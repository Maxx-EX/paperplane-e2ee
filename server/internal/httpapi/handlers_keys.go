package httpapi

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/paperplane/paperplane/server/internal/store"
)

var accountIDPattern = regexp.MustCompile(`^[a-f0-9]{64}$`)

type preKeyJSON struct {
	ID        int64  `json:"id"`
	Public    string `json:"public"`
	Signature string `json:"signature,omitempty"`
}

type oneTimePreKeyJSON struct {
	ID     int64  `json:"id"`
	Public string `json:"public"`
}

type bundleRequest struct {
	IdentityKey    string             `json:"identity_key"`
	SignedPreKey   preKeyJSON         `json:"signed_prekey"`
	KyberPreKey    *preKeyJSON        `json:"kyber_prekey,omitempty"`
	OneTimePreKeys []oneTimePreKeyJSON `json:"one_time_prekeys"`
}

type bundleResponse struct {
	AccountID     string              `json:"account_id"`
	IdentityKey   string              `json:"identity_key"`
	SignedPreKey  preKeyJSON          `json:"signed_prekey"`
	KyberPreKey   *preKeyJSON         `json:"kyber_prekey,omitempty"`
	OneTimePreKey *oneTimePreKeyJSON  `json:"one_time_prekey,omitempty"`
}

func (s *Server) handleUploadKeys(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	if !s.limiter.AllowTokenBucket(r.Context(), "prekey", accountID, s.cfg.Defense.RateLimits.PreKeyPerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "too many pre-key updates")
		return
	}
	var req bundleRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}

	identity, err := decodeKey(req.IdentityKey, 32, 128)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_identity_key", err.Error())
		return
	}
	signed, err := decodePreKey(req.SignedPreKey, true)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_signed_prekey", err.Error())
		return
	}
	if req.KyberPreKey != nil && req.KyberPreKey.Public != "" {
		if _, err := decodePreKey(*req.KyberPreKey, false); err != nil {
			writeErr(w, http.StatusBadRequest, "invalid_kyber_prekey", err.Error())
			return
		}
	}
	if len(req.OneTimePreKeys) == 0 || len(req.OneTimePreKeys) > 200 {
		writeErr(w, http.StatusBadRequest, "invalid_one_time_prekeys", "must upload 1-200 one-time prekeys")
		return
	}

	otps := make([]store.OneTimePreKey, 0, len(req.OneTimePreKeys))
	for _, k := range req.OneTimePreKeys {
		if k.ID <= 0 {
			writeErr(w, http.StatusBadRequest, "invalid_one_time_prekeys", "prekey id must be positive")
			return
		}
		pub, err := decodeKey(k.Public, 32, 128)
		if err != nil {
			writeErr(w, http.StatusBadRequest, "invalid_one_time_prekeys", err.Error())
			return
		}
		otps = append(otps, store.OneTimePreKey{ID: k.ID, PublicKey: pub})
	}

	bundle := store.Bundle{
		AccountID:             accountID,
		IdentityKey:           identity,
		SignedPreKeyID:        signed.ID,
		SignedPreKeyPublic:    signed.PublicKey,
		SignedPreKeySignature: signed.Signature,
	}
	if req.KyberPreKey != nil && req.KyberPreKey.Public != "" {
		kyber, _ := decodePreKey(*req.KyberPreKey, false)
		bundle.KyberPreKeyID = kyber.ID
		bundle.KyberPreKeyPublic = kyber.PublicKey
		bundle.KyberPreKeySignature = kyber.Signature
	}

	ctx, cancel := timeout(r, 5*time.Second)
	defer cancel()
	if err := s.db.UpsertBundle(ctx, bundle, otps); err != nil {
		s.log.Error("bundle upload failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"one_time_prekeys_stored": len(otps),
		"identity_key":            req.IdentityKey,
	})
}

func (s *Server) handleGetBundle(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	if !s.limiter.AllowTokenBucket(r.Context(), "prekey", accountID, s.cfg.Defense.RateLimits.PreKeyPerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "too many pre-key fetches")
		return
	}
	target := r.PathValue("accountID")
	if !accountIDPattern.MatchString(target) {
		writeErr(w, http.StatusBadRequest, "invalid_account", "account id must be 64 lowercase hex chars")
		return
	}
	ctx, cancel := timeout(r, 5*time.Second)
	defer cancel()

	b, err := s.db.TakeBundle(ctx, target)
	if errors.Is(err, store.ErrNotFound) {
		writeErr(w, http.StatusNotFound, "no_bundle", "recipient has not published keys")
		return
	}
	if err != nil {
		s.log.Error("bundle fetch failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}

	resp := bundleResponse{
		AccountID:   b.AccountID,
		IdentityKey: b64(b.IdentityKey),
		SignedPreKey: preKeyJSON{
			ID: b.SignedPreKeyID, Public: b64(b.SignedPreKeyPublic),
			Signature: b64(b.SignedPreKeySignature),
		},
	}
	if len(b.KyberPreKeyPublic) > 0 {
		resp.KyberPreKey = &preKeyJSON{
			ID: b.KyberPreKeyID, Public: b64(b.KyberPreKeyPublic),
			Signature: b64(b.KyberPreKeySignature),
		}
	}
	if b.OneTimePreKey != nil {
		resp.OneTimePreKey = &oneTimePreKeyJSON{
			ID: b.OneTimePreKey.ID, Public: b64(b.OneTimePreKey.PublicKey),
		}
	}
	writeJSON(w, http.StatusOK, resp)
}

// ------------------------------------------------------------------ decoding

func b64(b []byte) string { return base64.StdEncoding.EncodeToString(b) }

func decodeKey(s string, minLen, maxLen int) ([]byte, error) {
	raw, err := decodeB64(s)
	if err != nil {
		return nil, errors.New("value is not valid base64")
	}
	if len(raw) < minLen || len(raw) > maxLen {
		return nil, fmt.Errorf("key length %d outside %d-%d bytes", len(raw), minLen, maxLen)
	}
	return raw, nil
}

func decodePreKey(k preKeyJSON, requireSignature bool) (struct {
	ID        int64
	PublicKey []byte
	Signature []byte
}, error) {
	var out struct {
		ID        int64
		PublicKey []byte
		Signature []byte
	}
	if k.ID <= 0 {
		return out, errors.New("prekey id must be positive")
	}
	pub, err := decodeKey(k.Public, 32, 128)
	if err != nil {
		return out, err
	}
	out.ID, out.PublicKey = k.ID, pub
	if k.Signature != "" {
		sig, err := decodeKey(k.Signature, 32, 128)
		if err != nil {
			return out, err
		}
		out.Signature = sig
	} else if requireSignature {
		return out, errors.New("signed prekey requires a signature")
	}
	return out, nil
}

func decodeB64(s string) ([]byte, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("empty")
	}
	for _, enc := range []*base64.Encoding{
		base64.StdEncoding, base64.RawStdEncoding, base64.URLEncoding, base64.RawURLEncoding,
	} {
		if raw, err := enc.DecodeString(s); err == nil {
			return raw, nil
		}
	}
	return nil, errors.New("invalid base64")
}
