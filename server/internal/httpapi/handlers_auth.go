package httpapi

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/paperplane/paperplane/server/internal/auth"
	"github.com/paperplane/paperplane/server/internal/store"
)

const (
	nonceTTL    = 2 * time.Minute
	ticketTTL   = 30 * time.Second
	sealedTTL   = 24 * time.Hour
	sealedBatch = 5
)

type challengeRequest struct {
	Username string `json:"username"`
	Intent   string `json:"intent"` // "register" | "login"
}

type challengeResponse struct {
	AccountID string `json:"account_id"`
	Nonce     string `json:"nonce"`
	ExpiresIn int    `json:"expires_in"`
}

// handleChallenge issues a single-use nonce bound to the derived account id.
// For login we deliberately answer identically whether or not the account
// exists, so the endpoint cannot be used as a username oracle.
func (s *Server) handleChallenge(w http.ResponseWriter, r *http.Request) {
	ip := s.clientIP(r)
	if !s.limiter.AllowTokenBucket(r.Context(), "challenge", ip, s.cfg.Defense.RateLimits.ChallengePerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "too many challenges")
		return
	}
	var req challengeRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}
	username := auth.NormalizeUsername(req.Username)
	if !auth.ValidUsername(username) {
		writeErr(w, http.StatusBadRequest, "invalid_username", "username must be 3-32 chars: a-z 0-9 . _ -")
		return
	}
	accountID := auth.AccountID(s.cfg.Crypto.Pepper, username)

	ctx, cancel := timeout(r, 3*time.Second)
	defer cancel()

	switch strings.ToLower(strings.TrimSpace(req.Intent)) {
	case "register":
		exists, err := s.db.AccountExists(ctx, accountID)
		if err != nil {
			s.log.Error("account lookup failed", "err", err)
			writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
			return
		}
		if exists {
			writeErr(w, http.StatusConflict, "account_exists", "username already registered")
			return
		}
	case "login", "":
	default:
		writeErr(w, http.StatusBadRequest, "invalid_intent", "intent must be login or register")
		return
	}

	nonce, err := auth.NewNonce()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "rng_unavailable", "could not issue nonce")
		return
	}
	if err := s.rdb.Set(ctx, nonceKey(accountID, nonce), "1", nonceTTL).Err(); err != nil {
		s.log.Error("nonce store failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	writeJSON(w, http.StatusOK, challengeResponse{
		AccountID: accountID, Nonce: nonce, ExpiresIn: int(nonceTTL.Seconds()),
	})
}

type registerRequest struct {
	Username     string `json:"username"`
	Argon2IDHash string `json:"argon2id_hash"`
	Nonce        string `json:"nonce"`
	Proof        string `json:"proof"`
}

type sessionResponse struct {
	Token     string `json:"token"`
	AccountID string `json:"account_id"`
	ExpiresIn int    `json:"expires_in"`
}

// handleRegister stores only the PHC string. The client-proved nonce ensures
// the caller actually holds the credential it is asking us to bind.
func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	ip := s.clientIP(r)
	if !s.limiter.AllowTokenBucket(r.Context(), "register", ip, s.cfg.Defense.RateLimits.ChallengePerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "too many registrations")
		return
	}
	var req registerRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}
	username := auth.NormalizeUsername(req.Username)
	if !auth.ValidUsername(username) {
		writeErr(w, http.StatusBadRequest, "invalid_username", "invalid username")
		return
	}
	accountID := auth.AccountID(s.cfg.Crypto.Pepper, username)

	mem, iter, par, err := auth.ParsePHC(req.Argon2IDHash)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_credential", err.Error())
		return
	}
	// Refuse downgraded parameters: a client must not register with cheaper
	// Argon2 settings than the server advertises.
	if mem < s.cfg.Crypto.Argon2.MemoryKiB || iter < s.cfg.Crypto.Argon2.Iterations ||
		par < s.cfg.Crypto.Argon2.Parallelism {
		writeErr(w, http.StatusBadRequest, "weak_credential", "argon2id parameters below server policy")
		return
	}

	ctx, cancel := timeout(r, 8*time.Second)
	defer cancel()

	if !s.consumeNonce(ctx, accountID, req.Nonce) {
		writeErr(w, http.StatusBadRequest, "invalid_nonce", "nonce missing, reused or expired")
		return
	}
	if !auth.VerifyProof(req.Argon2IDHash, req.Nonce, req.Proof) {
		writeErr(w, http.StatusUnauthorized, "invalid_proof", "proof does not match credential")
		return
	}

	if err := s.db.CreateAccount(ctx, accountID, req.Argon2IDHash); err != nil {
		if errors.Is(err, store.ErrAccountExists) {
			writeErr(w, http.StatusConflict, "account_exists", "username already registered")
			return
		}
		s.log.Error("create account failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}

	token, err := s.jwt.Issue(accountID)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "token_error", "could not issue token")
		return
	}
	s.limiter.PublishEvent(r.Context(), "register", map[string]string{"account": accountID, "ip": ip})
	writeJSON(w, http.StatusCreated, sessionResponse{
		Token: token, AccountID: accountID, ExpiresIn: int(s.jwt.TTL().Seconds()),
	})
}

type loginRequest struct {
	Username string `json:"username"`
	Nonce    string `json:"nonce"`
	Proof    string `json:"proof"`
}

// handleLogin never receives the credential: the client proves knowledge of
// the stored hash by HMAC-ing the nonce. A wrong password simply fails the
// constant-time proof check.
func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	ip := s.clientIP(r)
	if !s.limiter.AllowTokenBucket(r.Context(), "login", ip, s.cfg.Defense.RateLimits.ChallengePerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "too many attempts")
		return
	}
	var req loginRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}
	username := auth.NormalizeUsername(req.Username)
	if !auth.ValidUsername(username) {
		writeErr(w, http.StatusUnauthorized, "invalid_credentials", "invalid credentials")
		return
	}
	accountID := auth.AccountID(s.cfg.Crypto.Pepper, username)

	ctx, cancel := timeout(r, 8*time.Second)
	defer cancel()

	used := s.consumeNonce(ctx, accountID, req.Nonce)
	stored, err := s.db.AuthHash(ctx, accountID)
	if err != nil && !errors.Is(err, store.ErrNotFound) {
		s.log.Error("auth hash lookup failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	// Burn the nonce even when the account is unknown, and keep the response
	// identical for both failure paths.
	if !used || errors.Is(err, store.ErrNotFound) ||
		!auth.VerifyProof(stored, req.Nonce, req.Proof) {
		writeErr(w, http.StatusUnauthorized, "invalid_credentials", "invalid credentials")
		return
	}

	token, err := s.jwt.Issue(accountID)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "token_error", "could not issue token")
		return
	}
	s.db.TouchAccount(ctx, accountID)
	writeJSON(w, http.StatusOK, sessionResponse{
		Token: token, AccountID: accountID, ExpiresIn: int(s.jwt.TTL().Seconds()),
	})
}

func (s *Server) consumeNonce(ctx context.Context, accountID, nonce string) bool {
	if nonce == "" || len(nonce) > 64 {
		return false
	}
	val, err := s.rdb.GetDel(ctx, nonceKey(accountID, nonce)).Result()
	if errors.Is(err, redis.Nil) {
		return false
	}
	return err == nil && val != ""
}

func nonceKey(accountID, nonce string) string {
	return "paperplane:nonce:" + accountID + ":" + nonce
}

// handleWSTicket trades a long-lived JWT for a 30-second single-use ticket.
// Tokens therefore never appear in URLs or proxy access logs.
func (s *Server) handleWSTicket(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	ticket, err := auth.NewTicket()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "rng_unavailable", "could not issue ticket")
		return
	}
	ctx, cancel := timeout(r, 3*time.Second)
	defer cancel()
	if err := s.rdb.Set(ctx, "paperplane:wsticket:"+ticket, accountID, ticketTTL).Err(); err != nil {
		s.log.Error("ticket store failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ticket": ticket, "expires_in": int(ticketTTL.Seconds())})
}

// handleSealedTokens mints blind sender tokens. Each token can be attached to
// exactly one envelope, after which the server stores no sender at all.
func (s *Server) handleSealedTokens(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	ctx, cancel := timeout(r, 3*time.Second)
	defer cancel()

	tokens := make([]string, 0, sealedBatch)
	for i := 0; i < sealedBatch; i++ {
		t, err := auth.NewTicket()
		if err != nil {
			writeErr(w, http.StatusInternalServerError, "rng_unavailable", "could not mint token")
			return
		}
		if err := s.rdb.Set(ctx, "paperplane:sealed:"+t, accountID, sealedTTL).Err(); err != nil {
			s.log.Error("sealed token store failed", "err", err)
			writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
			return
		}
		tokens = append(tokens, t)
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"tokens": tokens, "expires_in": int(sealedTTL.Seconds()),
	})
}
