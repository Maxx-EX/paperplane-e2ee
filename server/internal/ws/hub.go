// Package ws owns the WebSocket fan-out. One account may have several live
// sockets (phone + laptop); each connection gets its own bounded send queue so
// a single slow reader can never block delivery to everyone else.
package ws

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 90 * time.Second
	pingPeriod     = 30 * time.Second
	maxMessageSize = 64 << 10
	sendQueueSize  = 32
)

type Client struct {
	hub       *Hub
	conn      *websocket.Conn
	AccountID string
	IP        string

	send      chan []byte
	done      chan struct{}
	closeOnce sync.Once
}

type Hub struct {
	mu        sync.RWMutex
	byAccount map[string]map[*Client]struct{}
	byIP      map[string]int
	log       *slog.Logger
	closed    bool
}

func NewHub(log *slog.Logger) *Hub {
	return &Hub{
		byAccount: map[string]map[*Client]struct{}{},
		byIP:      map[string]int{},
		log:       log,
	}
}

// Add registers a freshly upgraded connection and returns the client handle.
func (h *Hub) Add(accountID, ip string, conn *websocket.Conn) *Client {
	c := &Client{
		hub:       h,
		conn:      conn,
		AccountID: accountID,
		IP:        ip,
		send:      make(chan []byte, sendQueueSize),
		done:      make(chan struct{}),
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.byAccount[accountID] == nil {
		h.byAccount[accountID] = map[*Client]struct{}{}
	}
	h.byAccount[accountID][c] = struct{}{}
	h.byIP[ip]++
	return c
}

func (h *Hub) remove(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if set, ok := h.byAccount[c.AccountID]; ok {
		delete(set, c)
		if len(set) == 0 {
			delete(h.byAccount, c.AccountID)
		}
	}
	if c.IP != "" {
		h.byIP[c.IP]--
		if h.byIP[c.IP] <= 0 {
			delete(h.byIP, c.IP)
		}
	}
}

// ConnectionsForIP is used by the handshake limiter (MAX_WS_PER_IP).
func (h *Hub) ConnectionsForIP(ip string) int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.byIP[ip]
}

// Deliver pushes a frame to every live socket of an account. It reports
// whether at least one socket accepted the frame; otherwise the caller keeps
// the envelope in PostgreSQL for the offline pull path.
func (h *Hub) Deliver(accountID string, payload []byte) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	delivered := false
	for c := range h.byAccount[accountID] {
		select {
		case c.send <- payload:
			delivered = true
		default:
			h.log.Warn("slow websocket client, dropping push", "account", shortID(accountID))
		}
	}
	return delivered
}

func (h *Hub) CloseAll() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.closed = true
	for _, set := range h.byAccount {
		for c := range set {
			c.Close()
		}
	}
}

func (c *Client) Start(ctx context.Context) {
	go c.writePump(ctx)
}

func (c *Client) writePump(ctx context.Context) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.done:
			return
		case msg := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Send queues a frame without blocking. Returns false when the queue is full.
func (c *Client) Send(payload []byte) bool {
	select {
	case c.send <- payload:
		return true
	default:
		return false
	}
}

func (c *Client) Close() {
	c.closeOnce.Do(func() {
		close(c.done)
		c.hub.remove(c)
		_ = c.conn.WriteControl(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			time.Now().Add(writeWait))
		_ = c.conn.Close()
	})
}

// ------------------------------------------------------------ frame schema

// EnvelopeFrame is the only message type the server proactively pushes.
type EnvelopeFrame struct {
	Type        string    `json:"type"`
	ID          int64     `json:"id"`
	From        string    `json:"from,omitempty"`
	CipherType  int       `json:"cipher_type"`
	Body        string    `json:"body"`
	ClientMsgID string    `json:"client_msg_id"`
	SentAt      time.Time `json:"sent_at"`
}

type ReadyFrame struct {
	Type      string `json:"type"`
	AccountID string `json:"account_id"`
	ServerTS  int64  `json:"server_ts"`
}

func shortID(id string) string {
	if len(id) <= 10 {
		return id
	}
	return id[:10] + "…"
}
