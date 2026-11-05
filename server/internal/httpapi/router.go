// Package httpapi is the only HTTP surface of the server. Every handler is
// deliberately small: validate, rate-limit, touch the store, return JSON.
// Ciphertext is opaque here — no handler inspects a message body.
package httpapi

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/paperplane/paperplane/server/internal/auth"
	"github.com/paperplane/paperplane/server/internal/config"
	"github.com/paperplane/paperplane/server/internal/defense"
	"github.com/paperplane/paperplane/server/internal/store"
	"github.com/paperplane/paperplane/server/internal/ws"
)

const maxBodyBytes = 2 << 20 // 2 MiB encoded; decoded ciphertext capped at 1 MiB

type Server struct {
	cfg     *config.Config
	db      *store.Postgres
	rdb     *redis.Client
	limiter *defense.Limiter
	hub     *ws.Hub
	jwt     *auth.Issuer
	log     *slog.Logger
	started time.Time
}

func New(cfg *config.Config, db *store.Postgres, rdb *redis.Client, lim *defense.Limiter,
	hub *ws.Hub, jwt *auth.Issuer, log *slog.Logger) *Server {
	return &Server{
		cfg: cfg, db: db, rdb: rdb, limiter: lim, hub: hub, jwt: jwt,
		log: log, started: time.Now(),
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", s.handleHealth)
	mux.HandleFunc("GET /api/v1/params", s.handleParams)

	mux.HandleFunc("POST /api/v1/auth/challenge", s.handleChallenge)
	mux.HandleFunc("POST /api/v1/auth/register", s.handleRegister)
	mux.HandleFunc("POST /api/v1/auth/login", s.handleLogin)

	mux.HandleFunc("POST /api/v1/ws-ticket", s.requireAuth(s.handleWSTicket))
	mux.HandleFunc("POST /api/v1/sealed-tokens", s.requireAuth(s.handleSealedTokens))

	mux.HandleFunc("POST /api/v1/keys", s.requireAuth(s.handleUploadKeys))
	mux.HandleFunc("GET /api/v1/keys/{accountID}", s.requireAuth(s.handleGetBundle))

	mux.HandleFunc("POST /api/v1/envelopes", s.requireAuth(s.handlePostEnvelope))
	mux.HandleFunc("GET /api/v1/envelopes", s.requireAuth(s.handlePullEnvelopes))
	mux.HandleFunc("POST /api/v1/envelopes/ack", s.requireAuth(s.handleAckEnvelopes))

	mux.HandleFunc("GET /ws", s.handleWS)
	mux.HandleFunc("/", s.handleNotFound)

	return s.common(mux)
}

// ------------------------------------------------------------------ helpers

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	if v == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Debug("write response", "err", err)
	}
}

func writeErr(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{
		"error": map[string]string{"code": code, "message": message},
	})
}

func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}

func (s *Server) clientIP(r *http.Request) string {
	// The service binds loopback and is reachable only through Caddy, which
	// sets X-Forwarded-For. Trust the left-most well-formed address.
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		first := strings.TrimSpace(strings.Split(xff, ",")[0])
		if ip := net.ParseIP(first); ip != nil {
			return first
		}
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func (s *Server) originAllowed(raw string) bool {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		// Native clients send no Origin; browsers always do.
		return true
	}
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	normalized := strings.ToLower(u.Scheme + "://" + u.Host)
	for _, allowed := range s.cfg.Server.WSAllowedOrigins {
		au, err := url.Parse(strings.TrimSpace(allowed))
		if err != nil || au.Scheme == "" || au.Host == "" {
			continue
		}
		if strings.ToLower(au.Scheme+"://"+au.Host) == normalized {
			return true
		}
	}
	return false
}

func jsonMarshal(v any) ([]byte, error) { return json.Marshal(v) }

func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	writeErr(w, http.StatusNotFound, "not_found", "unknown route")
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := timeout(r, 2*time.Second)
	defer cancel()
	status := "ok"
	if err := s.db.Ping(ctx); err != nil {
		status = "degraded"
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"status":   status,
		"uptime_s": int(time.Since(s.started).Seconds()),
		"time":     time.Now().UTC().Format(time.RFC3339),
	})
}

type paramsResponse struct {
	Project        string `json:"project"`
	PublicURL      string `json:"public_url"`
	Pepper         string `json:"pepper"`
	Argon2Memory   uint32 `json:"argon2_memory_kib"`
	Argon2Iter     uint32 `json:"argon2_iterations"`
	Argon2Parallel uint8  `json:"argon2_parallelism"`
	WSPath         string `json:"ws_path"`
	EnvelopeTTL    int    `json:"envelope_ttl_hours"`
	MaxEnvelope    int    `json:"max_envelope_bytes"`
	TokenTTL       int    `json:"token_ttl_seconds"`
}

// handleParams publishes exactly what a client needs to derive credentials.
// The pepper is a public domain separator by design: account ids stay
// irreversible because the server never stores a username.
func (s *Server) handleParams(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, paramsResponse{
		Project:        s.cfg.Project,
		PublicURL:      s.cfg.Server.PublicURL,
		Pepper:         s.cfg.Crypto.Pepper,
		Argon2Memory:   s.cfg.Crypto.Argon2.MemoryKiB,
		Argon2Iter:     s.cfg.Crypto.Argon2.Iterations,
		Argon2Parallel: s.cfg.Crypto.Argon2.Parallelism,
		WSPath:         "/ws",
		EnvelopeTTL:    s.cfg.Database.EnvelopeTTLHour,
		MaxEnvelope:    1 << 20,
		TokenTTL:       int(s.jwt.TTL().Seconds()),
	})
}

// unwrapBody normalises EOF (empty body) into a friendlier error.
func unwrapBody(err error) error {
	if errors.Is(err, io.EOF) {
		return errors.New("empty request body")
	}
	return err
}
