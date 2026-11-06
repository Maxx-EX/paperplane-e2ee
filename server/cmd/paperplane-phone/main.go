// paperplane-phone: 手机/Termux 专用极简 relay。
// 零外部依赖：不要 PostgreSQL、不要 Redis、不要 Docker、不要 proot。
// 纯内存存储，重启即清空，只给局域网内 1-2 个人测试用。
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type account struct {
	Hash   string            `json:"hash"`
	Bundle map[string]any    `json:"bundle"`
	WS     *websocket.Conn   `json:"-"`
	mu     sync.Mutex        `json:"-"`
}

type envelope struct {
	ID         int64  `json:"id"`
	From       string `json:"from"`
	To         string `json:"to"`
	CipherType int    `json:"cipher_type"`
	Body       string `json:"body"`
	SentAt     int64  `json:"sent_at"`
}

var (
	mu       sync.RWMutex
	accounts = map[string]*account{}
	outbox  = map[string][]envelope{}
	nextID  int64 = 1
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("/api/v1/params", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"pepper": "dev-pepper-change-me",
			"argon2": map[string]int{"m": 65536, "t": 3, "p": 1},
		})
	})
	mux.HandleFunc("/api/v1/auth/challenge", func(w http.ResponseWriter, r *http.Request) {
		var body struct{ Username, Action string }
		json.NewDecoder(r.Body).Decode(&body)
		json.NewEncoder(w).Encode(map[string]any{"nonce": fmt.Sprintf("nonce-%d", time.Now().UnixNano())})
	})
	mux.HandleFunc("/api/v1/auth/register", handleRegister)
	mux.HandleFunc("/api/v1/auth/login", handleLogin)
	mux.HandleFunc("/api/v1/keys", handleKeysUpload)
	mux.HandleFunc("/api/v1/keys/{id}", handleKeysGet)
	mux.HandleFunc("/api/v1/envelopes", handleEnvelopes)
	mux.HandleFunc("/api/v1/envelopes/ack", handleAck)
	mux.HandleFunc("/api/v1/ws-ticket", handleWsTicket)
	mux.HandleFunc("/ws", handleWS)
	// 根路径：返回版本号，方便用户在手机浏览器里访问 http://127.0.0.1:8080 验证
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			json.NewEncoder(w).Encode(map[string]any{
				"app": "paperplane-phone", "version": "v0.3.7", "status": "ok",
			})
			return
		}
		// 所有未匹配的路由返回 JSON，不要返回 Go 默认的文本 404（会让客户端 JSON 解析崩）
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{"error": "not found", "path": r.URL.Path})
	})

	srv := &http.Server{Addr: *addr, Handler: mux}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		slog.Info("paperplane-phone v0.3.7 listening", "addr", *addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server", "err", err)
		}
	}()
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	srv.Shutdown(shutdownCtx)
}

func authUser(r *http.Request) string {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if !strings.HasPrefix(token, "dev-") {
		return ""
	}
	return strings.TrimPrefix(token, "dev-")
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Hash     string `json:"argon2id_hash"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	mu.Lock()
	accounts[body.Username] = &account{Hash: body.Hash, Bundle: map[string]any{}}
	mu.Unlock()
	json.NewEncoder(w).Encode(map[string]any{
		"ok": true, "token": "dev-" + body.Username, "account_id": body.Username,
	})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	mu.RLock()
	_, ok := accounts[body.Username]
	mu.RUnlock()
	if !ok {
		http.Error(w, "no such user", 404)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"ok": true, "token": "dev-" + body.Username, "account_id": body.Username,
	})
}

func handleKeysUpload(w http.ResponseWriter, r *http.Request) {
	aid := authUser(r)
	if aid == "" {
		http.Error(w, "unauthorized", 401)
		return
	}
	var bundle map[string]any
	json.NewDecoder(r.Body).Decode(&bundle)
	mu.Lock()
	if acct, ok := accounts[aid]; ok {
		acct.Bundle = bundle
	}
	mu.Unlock()
	json.NewEncoder(w).Encode(map[string]any{"ok": true})
}

func handleKeysGet(w http.ResponseWriter, r *http.Request) {
	aid := r.PathValue("id")
	mu.RLock()
	acct, ok := accounts[aid]
	mu.RUnlock()
	if !ok {
		http.Error(w, "no such user", 404)
		return
	}
	json.NewEncoder(w).Encode(acct.Bundle)
}

func handleEnvelopes(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var body struct {
			To         string `json:"to"`
			CipherType int    `json:"cipher_type"`
			Body       string `json:"body"`
		}
		json.NewDecoder(r.Body).Decode(&body)
		from := authUser(r)
		mu.Lock()
		env := envelope{
			ID: nextID, From: from, To: body.To,
			CipherType: body.CipherType, Body: body.Body,
			SentAt: time.Now().UnixMilli(),
		}
		nextID++
		outbox[body.To] = append(outbox[body.To], env)
		if acct, ok := accounts[body.To]; ok && acct.WS != nil {
			acct.mu.Lock()
			c := acct.WS
			acct.mu.Unlock()
			_ = c.WriteJSON(env)
		}
		mu.Unlock()
		json.NewEncoder(w).Encode(map[string]any{"ok": true, "id": env.ID})
		return
	}
	// GET: pull pending envelopes
	aid := authUser(r)
	if aid == "" {
		http.Error(w, "unauthorized", 401)
		return
	}
	mu.RLock()
	pending := outbox[aid]
	mu.RUnlock()
	json.NewEncoder(w).Encode(map[string]any{"envelopes": pending})
}

func handleAck(w http.ResponseWriter, r *http.Request) {
	var body struct {
		IDs []int64 `json:"ids"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	aid := authUser(r)
	if aid == "" {
		http.Error(w, "unauthorized", 401)
		return
	}
	mu.Lock()
	keep := []envelope{}
	for _, e := range outbox[aid] {
		acked := false
		for _, id := range body.IDs {
			if e.ID == id {
				acked = true
				break
			}
		}
		if !acked {
			keep = append(keep, e)
		}
	}
	outbox[aid] = keep
	mu.Unlock()
	json.NewEncoder(w).Encode(map[string]any{"ok": true})
}

func handleWsTicket(w http.ResponseWriter, r *http.Request) {
	aid := authUser(r)
	if aid == "" {
		http.Error(w, "unauthorized", 401)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"ticket": "dev-" + aid})
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if !strings.HasPrefix(token, "dev-") {
		http.Error(w, "unauthorized", 401)
		return
	}
	aid := strings.TrimPrefix(token, "dev-")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	mu.Lock()
	if acct, ok := accounts[aid]; ok {
		acct.mu.Lock()
		acct.WS = c
		acct.mu.Unlock()
		for _, e := range outbox[aid] {
			_ = c.WriteJSON(e)
		}
		outbox[aid] = nil
	}
	mu.Unlock()
	defer c.Close()
	for {
		if _, _, err := c.ReadMessage(); err != nil {
			break
		}
	}
}
