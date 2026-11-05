// paperplane-phone: 手机/Termux 专用极简 relay。
// 零外部依赖：不要 PostgreSQL、不要 Redis、不要 Docker、不要 proot。
// 纯内存存储，重启即清空，只给局域网内 1-2 个人测试用。
// 生产部署请用 server/cmd/server（PostgreSQL + Redis + Caddy + CrowdSec）。
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
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // 局域网测试
}

type account struct {
	Argon2Hash string            `json:"argon2_hash"`
	Bundle     map[string]any    `json:"bundle"`
	WS         *websocket.Conn   `json:"-"`
	mu         sync.Mutex        `json:"-"`
}

type envelope struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Type    int    `json:"type"`
	Body    string `json:"body"`
	SentAt  int64  `json:"sent_at"`
}

var (
	mu       sync.RWMutex
	accounts = map[string]*account{}
	outbox  = map[string][]envelope{}
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("/v1/params", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"pepper": "dev-pepper-change-me",
			"argon2": map[string]int{"m": 65536, "t": 3, "p": 1},
		})
	})
	mux.HandleFunc("/v1/challenge", func(w http.ResponseWriter, r *http.Request) {
		var body struct{ Username, Action string }
		json.NewDecoder(r.Body).Decode(&body)
		json.NewEncoder(w).Encode(map[string]any{"nonce": fmt.Sprintf("nonce-%d", time.Now().UnixNano())})
	})
	mux.HandleFunc("/v1/register", handleRegister)
	mux.HandleFunc("/v1/login", handleLogin)
	mux.HandleFunc("/v1/keys/upload", handleKeysUpload)
	mux.HandleFunc("/v1/keys/:id", handleKeysGet)
	mux.HandleFunc("/v1/envelopes", handleEnvelope)
	mux.HandleFunc("/ws", handleWS)

	srv := &http.Server{Addr: *addr, Handler: mux}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		slog.Info("paperplane-phone listening", "addr", *addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server", "err", err)
		}
	}()
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	srv.Shutdown(shutdownCtx)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var body struct {
		AccountID  string `json:"account_id"`
		Argon2Hash string `json:"argon2_hash"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	mu.Lock()
	accounts[body.AccountID] = &account{Argon2Hash: body.Argon2Hash, Bundle: map[string]any{}}
	mu.Unlock()
	json.NewEncoder(w).Encode(map[string]any{"ok": true, "token": "dev-" + body.AccountID})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var body struct {
		AccountID string `json:"account_id"`
		Proof     string `json:"proof"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	mu.RLock()
	_, ok := accounts[body.AccountID]
	mu.RUnlock()
	if !ok {
		http.Error(w, "no such user", 404)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"ok": true, "token": "dev-" + body.AccountID})
}

func handleKeysUpload(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if len(token) < 5 || token[:4] != "dev-" {
		http.Error(w, "unauthorized", 401)
		return
	}
	aid := token[4:]
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

func handleEnvelope(w http.ResponseWriter, r *http.Request) {
	var env envelope
	json.NewDecoder(r.Body).Decode(&env)
	env.SentAt = time.Now().UnixMilli()
	mu.Lock()
	outbox[env.To] = append(outbox[env.To], env)
	// 如果对方在线，直接 WS 推送
	if acct, ok := accounts[env.To]; ok && acct.WS != nil {
		acct.mu.Lock()
		c := acct.WS
		acct.mu.Unlock()
		_ = c.WriteJSON(env)
	}
	mu.Unlock()
	json.NewEncoder(w).Encode(map[string]any{"ok": true})
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if len(token) < 5 || token[:4] != "dev-" {
		http.Error(w, "unauthorized", 401)
		return
	}
	aid := token[4:]
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	mu.Lock()
	if acct, ok := accounts[aid]; ok {
		acct.mu.Lock()
		acct.WS = c
		acct.mu.Unlock()
		// 推送离线信封
		for _, env := range outbox[aid] {
			_ = c.WriteJSON(env)
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
