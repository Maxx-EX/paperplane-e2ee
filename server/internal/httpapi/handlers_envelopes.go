package httpapi

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/gorilla/websocket"

	"github.com/paperplane/paperplane/server/internal/store"
	"github.com/paperplane/paperplane/server/internal/ws"
)

const maxEnvelopeBytes = 1 << 20 // decoded ciphertext cap

type envelopeRequest struct {
	To           string `json:"to"`
	ClientMsgID  string `json:"client_msg_id"`
	CipherType   int    `json:"cipher_type"`
	Body         string `json:"body"`
	SealedSender bool   `json:"sealed_sender"`
	SenderToken  string `json:"sender_token"`
}

type envelopeOut struct {
	ID          int64     `json:"id"`
	From        string    `json:"from,omitempty"`
	CipherType  int       `json:"cipher_type"`
	Body        string    `json:"body"`
	ClientMsgID string    `json:"client_msg_id"`
	SentAt      time.Time `json:"sent_at"`
}

func toOut(e store.Envelope) envelopeOut {
	return envelopeOut{
		ID: e.ID, From: e.SenderID, CipherType: e.CipherType,
		Body: base64.StdEncoding.EncodeToString(e.Body),
		ClientMsgID: e.ClientMsgID, SentAt: e.SentAt,
	}
}

func envelopeFrame(e store.Envelope) ws.EnvelopeFrame {
	return ws.EnvelopeFrame{
		Type: "envelope", ID: e.ID, From: e.SenderID, CipherType: e.CipherType,
		Body: base64.StdEncoding.EncodeToString(e.Body),
		ClientMsgID: e.ClientMsgID, SentAt: e.SentAt,
	}
}

// handlePostEnvelope accepts an opaque ciphertext. The server cannot decrypt
// it, does not log it, and deletes it once the recipient acknowledges delivery.
func (s *Server) handlePostEnvelope(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	ip := s.clientIP(r)

	if !s.limiter.AllowTokenBucket(r.Context(), "envelope", accountID, s.cfg.Defense.RateLimits.EnvelopePerMin) {
		writeErr(w, http.StatusTooManyRequests, "rate_limited", "envelope rate exceeded")
		return
	}

	var req envelopeRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}
	if !accountIDPattern.MatchString(req.To) {
		writeErr(w, http.StatusBadRequest, "invalid_recipient", "recipient must be a 64-char account id")
		return
	}
	if req.To == accountID {
		writeErr(w, http.StatusBadRequest, "invalid_recipient", "cannot send to self")
		return
	}
	if len(req.ClientMsgID) == 0 || len(req.ClientMsgID) > 80 {
		writeErr(w, http.StatusBadRequest, "invalid_client_msg_id", "client_msg_id must be 1-80 chars")
		return
	}
	if req.CipherType != 3 && req.CipherType != 4 {
		// 3 = SignalMessage, 4 = PreKeySignalMessage (libsignal wire types).
		writeErr(w, http.StatusBadRequest, "invalid_cipher_type", "cipher_type must be 3 or 4")
		return
	}
	body, err := base64.StdEncoding.DecodeString(req.Body)
	if err != nil {
		body, err = base64.RawStdEncoding.DecodeString(req.Body)
	}
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body_encoding", "ciphertext must be base64")
		return
	}
	if len(body) == 0 || len(body) > maxEnvelopeBytes {
		writeErr(w, http.StatusRequestEntityTooLarge, "envelope_too_large", "ciphertext must be 1 byte - 1 MiB")
		return
	}

	ctx, cancel := timeout(r, 8*time.Second)
	defer cancel()

	// Fan-out guard: one sender talking to too many distinct recipients inside
	// a window is blocked before any database write happens.
	if !s.limiter.AllowDistinctTargets(ctx, accountID, req.To,
		s.cfg.Defense.RateLimits.EnvelopeDistinctTargets, s.cfg.EnvelopeWindow()) {
		writeErr(w, http.StatusTooManyRequests, "fanout_limited", "too many distinct recipients")
		return
	}

	sender := accountID
	if req.SealedSender {
		if req.SenderToken == "" {
			writeErr(w, http.StatusBadRequest, "missing_sender_token", "sealed sender requires a token")
			return
		}
		owner, err := s.rdb.GetDel(ctx, "paperplane:sealed:"+req.SenderToken).Result()
		if errors.Is(err, redis.Nil) || (err == nil && owner != accountID) {
			writeErr(w, http.StatusUnauthorized, "invalid_sender_token", "sealed sender token rejected")
			return
		}
		if err != nil && !errors.Is(err, redis.Nil) {
			s.log.Error("sealed token lookup failed", "err", err)
			writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
			return
		}
		sender = "" // route metadata intentionally not persisted
	}

	exists, err := s.db.AccountExists(ctx, req.To)
	if err != nil {
		s.log.Error("recipient lookup failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	if !exists {
		writeErr(w, http.StatusNotFound, "no_recipient", "recipient has no account")
		return
	}

	stored := store.Envelope{
		TargetID: req.To, SenderID: sender, CipherType: req.CipherType,
		Body: body, ClientMsgID: req.ClientMsgID,
	}
	id, err := s.db.InsertEnvelope(ctx, stored, s.cfg.Database.EnvelopeTTLHour)
	if err != nil {
		s.log.Error("envelope insert failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	stored.ID = id
	stored.SentAt = time.Now().UTC()

	payload, _ := jsonMarshal(envelopeFrame(stored))
	delivered := s.hub.Deliver(req.To, payload)
	if delivered {
		// Once handed to a live socket, drop sender metadata immediately.
		s.db.StripSender(ctx, []int64{id})
	}
	s.limiter.PublishEvent(ctx, "envelope", map[string]string{
		"account": accountID, "ip": ip, "target": req.To,
		"sealed": strconv.FormatBool(req.SealedSender),
	})

	writeJSON(w, http.StatusCreated, map[string]any{
		"id": id, "delivered": delivered,
	})
}

func (s *Server) handlePullEnvelopes(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	limit := 50
	if raw := r.URL.Query().Get("limit"); raw != "" {
		if n, err := strconv.Atoi(raw); err == nil && n > 0 && n <= 100 {
			limit = n
		}
	}
	ctx, cancel := timeout(r, 8*time.Second)
	defer cancel()

	list, err := s.db.PendingEnvelopes(ctx, accountID, limit)
	if err != nil {
		s.log.Error("envelope pull failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	out := make([]envelopeOut, 0, len(list))
	for _, e := range list {
		out = append(out, toOut(e))
	}
	writeJSON(w, http.StatusOK, map[string]any{"envelopes": out, "count": len(out)})
}

type ackRequest struct {
	IDs []int64 `json:"ids"`
}

func (s *Server) handleAckEnvelopes(w http.ResponseWriter, r *http.Request) {
	accountID := accountFrom(r)
	var req ackRequest
	if err := decodeJSON(w, r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid_body", unwrapBody(err).Error())
		return
	}
	if len(req.IDs) == 0 || len(req.IDs) > 200 {
		writeErr(w, http.StatusBadRequest, "invalid_ids", "ids must contain 1-200 entries")
		return
	}
	for _, id := range req.IDs {
		if id <= 0 {
			writeErr(w, http.StatusBadRequest, "invalid_ids", "ids must be positive")
			return
		}
	}
	ctx, cancel := timeout(r, 5*time.Second)
	defer cancel()
	n, err := s.db.AckEnvelopes(ctx, accountID, req.IDs)
	if err != nil {
		s.log.Error("envelope ack failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"acked": n})
}

// ------------------------------------------------------------- websocket

func (s *Server) handleWS(w http.ResponseWriter, r *http.Request) {
	ip := s.clientIP(r)
	if !s.originAllowed(r.Header.Get("Origin")) {
		writeErr(w, http.StatusForbidden, "origin_denied", "websocket origin not allowed")
		return
	}
	if s.limiter.IsBanned(r.Context(), ip) {
		writeErr(w, http.StatusForbidden, "banned", "temporarily blocked")
		return
	}
	if s.hub.ConnectionsForIP(ip) >= s.cfg.Defense.RateLimits.WSPerIP {
		writeErr(w, http.StatusTooManyRequests, "ws_limit", "too many concurrent sockets from this address")
		return
	}

	ticket := r.URL.Query().Get("ticket")
	if ticket == "" || len(ticket) > 128 {
		writeErr(w, http.StatusUnauthorized, "missing_ticket", "single-use ticket required")
		return
	}
	ctx, cancel := timeout(r, 5*time.Second)
	accountID, err := s.rdb.GetDel(ctx, "paperplane:wsticket:"+ticket).Result()
	cancel()
	if errors.Is(err, redis.Nil) {
		writeErr(w, http.StatusUnauthorized, "invalid_ticket", "ticket expired or already used")
		return
	}
	if err != nil {
		s.log.Error("ticket lookup failed", "err", err)
		writeErr(w, http.StatusServiceUnavailable, "store_unavailable", "try again")
		return
	}
	if !accountIDPattern.MatchString(accountID) {
		writeErr(w, http.StatusUnauthorized, "invalid_ticket", "ticket rejected")
		return
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Exact whitelist only: never `return true`, never compare r.Host
		// (Caddy rewrites Host but preserves Origin).
		CheckOrigin: func(req *http.Request) bool {
			return s.originAllowed(req.Header.Get("Origin"))
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.log.Warn("websocket upgrade failed", "err", err)
		return
	}

	client := s.hub.Add(accountID, ip, conn)
	client.Start(r.Context())
	defer client.Close()

	conn.SetReadLimit(64 << 10)
	_ = conn.SetReadDeadline(time.Now().Add(90 * time.Second))
	conn.SetPongHandler(func(string) error {
		return conn.SetReadDeadline(time.Now().Add(90 * time.Second))
	})

	ready, _ := jsonMarshal(ws.ReadyFrame{
		Type: "ready", AccountID: accountID, ServerTS: time.Now().Unix(),
	})
	client.Send(ready)

	// Catch-up: deliver everything queued while the device was offline.
	pullCtx, pullCancel := timeout(r, 8*time.Second)
	pending, err := s.db.PendingEnvelopes(pullCtx, accountID, 50)
	pullCancel()
	if err != nil {
		s.log.Warn("pending pull on connect failed", "err", err)
	}
	pushed := make([]int64, 0, len(pending))
	for _, e := range pending {
		payload, _ := jsonMarshal(envelopeFrame(e))
		if client.Send(payload) {
			pushed = append(pushed, e.ID)
		}
	}
	if len(pushed) > 0 {
		stripCtx, stripCancel := timeout(r, 3*time.Second)
		s.db.StripSender(stripCtx, pushed)
		stripCancel()
	}

	s.limiter.PublishEvent(r.Context(), "connect", map[string]string{
		"account": accountID, "ip": ip,
		"reconnects": strconv.FormatInt(s.limiter.CountReconnect(r.Context(), accountID), 10),
	})

	var inbound struct {
		Type string  `json:"type"`
		IDs  []int64 `json:"ids"`
	}
	for {
		if err := conn.ReadJSON(&inbound); err != nil {
			return
		}
		switch inbound.Type {
		case "ack":
			if len(inbound.IDs) == 0 || len(inbound.IDs) > 200 {
				continue
			}
			ackCtx, ackCancel := timeout(r, 5*time.Second)
			if _, err := s.db.AckEnvelopes(ackCtx, accountID, inbound.IDs); err != nil {
				s.log.Warn("ws ack failed", "err", err)
			}
			ackCancel()
		case "ping":
			client.Send([]byte(`{"type":"pong"}`))
		}
	}
}
