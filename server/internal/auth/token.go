// Package auth implements the zero-knowledge account model and HS256 tokens.
//
// The server never sees a password. The client derives:
//
//	account_id = SHA-256(username || pepper)
//	argon2id   = Argon2id(password, salt = account_id[:16], m=64MiB, t=3, p=1)
//
// and sends only account_id (implicitly, by username) plus the PHC string.
// At login the client proves possession of that hash by HMAC-ing a single-use
// server nonce, so a captured credential cannot be replayed.
package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

const phcPrefix = "$argon2id$v=19$"

// base64url without padding keeps tokens URL/query safe.
var b64 = base64.RawURLEncoding

// UsernamePattern is deliberately small and ASCII-only: account ids are
// derived from it, so ambiguity would create duplicate-looking identities.
var UsernamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9._-]{2,31}$`)

// NormalizeUsername lowercases and trims; the result is what gets hashed.
func NormalizeUsername(raw string) string {
	return strings.ToLower(strings.TrimSpace(raw))
}

func ValidUsername(normalized string) bool { return UsernamePattern.MatchString(normalized) }

// AccountID is the irreversible account handle stored by the server.
func AccountID(pepper, normalizedUsername string) string {
	sum := sha256.Sum256([]byte(normalizedUsername + pepper))
	return hex.EncodeToString(sum[:])
}

// NewNonce returns 32 random bytes, base64url encoded.
func NewNonce() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return b64.EncodeToString(b), nil
}

// NewTicket returns a single-use WebSocket ticket.
func NewTicket() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return b64.EncodeToString(b), nil
}

// Proof binds the stored credential to a server nonce.
func Proof(credential, nonce string) string {
	mac := hmac.New(sha256.New, []byte(credential))
	mac.Write([]byte(nonce))
	return b64.EncodeToString(mac.Sum(nil))
}

func VerifyProof(credential, nonce, proof string) bool {
	expected := Proof(credential, nonce)
	return subtle.ConstantTimeCompare([]byte(expected), []byte(proof)) == 1
}

// ParsePHC validates the Argon2id PHC string shape and extracts m, t, p.
// The server rejects credentials weaker than the configured parameters.
func ParsePHC(encoded string) (memoryKiB, iterations uint32, parallelism uint8, err error) {
	if !strings.HasPrefix(encoded, phcPrefix) {
		return 0, 0, 0, errors.New("credential must be an argon2id v=19 PHC string")
	}
	rest := strings.TrimPrefix(encoded, phcPrefix)
	parts := strings.Split(rest, "$")
	if len(parts) != 3 {
		return 0, 0, 0, errors.New("malformed PHC string")
	}
	for _, kv := range strings.Split(parts[0], ",") {
		k, v, ok := strings.Cut(kv, "=")
		if !ok {
			return 0, 0, 0, errors.New("malformed argon2 parameters")
		}
		n, convErr := strconv.ParseUint(v, 10, 32)
		if convErr != nil {
			return 0, 0, 0, fmt.Errorf("bad argon2 parameter %q", kv)
		}
		switch k {
		case "m":
			memoryKiB = uint32(n)
		case "t":
			iterations = uint32(n)
		case "p":
			if n > 255 {
				return 0, 0, 0, errors.New("argon2 parallelism out of range")
			}
			parallelism = uint8(n)
		}
	}
	if memoryKiB == 0 || iterations == 0 || parallelism == 0 {
		return 0, 0, 0, errors.New("incomplete argon2 parameters")
	}
	if len(parts[1]) < 8 || len(parts[2]) < 16 {
		return 0, 0, 0, errors.New("argon2 salt/hash too short")
	}
	return memoryKiB, iterations, parallelism, nil
}

// Issuer creates and verifies HS256 JSON Web Tokens containing only account_id.
type Issuer struct {
	secret []byte
	ttl    time.Duration
	issuer string
	now    func() time.Time
}

type claims struct {
	Sub string `json:"sub"`
	Iss string `json:"iss"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
}

func NewIssuer(secret string, ttl time.Duration, issuer string) *Issuer {
	return &Issuer{secret: []byte(secret), ttl: ttl, issuer: issuer, now: time.Now}
}

// Issue returns a signed token. Header is fixed: {"alg":"HS256","typ":"JWT"}.
func (i *Issuer) Issue(accountID string) (string, error) {
	now := i.now()
	c := claims{Sub: accountID, Iss: i.issuer, Iat: now.Unix(), Exp: now.Add(i.ttl).Unix()}
	payload, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	header := b64.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	body := header + "." + b64.EncodeToString(payload)
	sig := i.sign(body)
	return body + "." + b64.EncodeToString(sig), nil
}

func (i *Issuer) sign(body string) []byte {
	mac := hmac.New(sha256.New, i.secret)
	mac.Write([]byte(body))
	return mac.Sum(nil)
}

// Verify checks structure, algorithm, signature, issuer and expiry.
func (i *Issuer) Verify(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", ErrInvalidToken
	}
	headerJSON, err := b64.DecodeString(parts[0])
	if err != nil {
		return "", ErrInvalidToken
	}
	var h struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}
	if err := json.Unmarshal(headerJSON, &h); err != nil || h.Alg != "HS256" {
		return "", ErrInvalidToken
	}
	expected := i.sign(parts[0] + "." + parts[1])
	got, err := b64.DecodeString(parts[2])
	if err != nil || subtle.ConstantTimeCompare(expected, got) != 1 {
		return "", ErrInvalidToken
	}
	payload, err := b64.DecodeString(parts[1])
	if err != nil {
		return "", ErrInvalidToken
	}
	var c claims
	if err := json.Unmarshal(payload, &c); err != nil {
		return "", ErrInvalidToken
	}
	if c.Iss != i.issuer || c.Sub == "" {
		return "", ErrInvalidToken
	}
	if i.now().Unix() >= c.Exp {
		return "", ErrExpiredToken
	}
	return c.Sub, nil
}

// TTL exposes the configured lifetime for API responses.
func (i *Issuer) TTL() time.Duration { return i.ttl }
