// Command server is the Paperplane relay. It terminates no TLS itself (Caddy
// does), decrypts nothing, and keeps ciphertext only until delivery is
// acknowledged.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/paperplane/paperplane/server/internal/auth"
	"github.com/paperplane/paperplane/server/internal/config"
	"github.com/paperplane/paperplane/server/internal/defense"
	"github.com/paperplane/paperplane/server/internal/httpapi"
	"github.com/paperplane/paperplane/server/internal/store"
	"github.com/paperplane/paperplane/server/internal/ws"
)

func main() {
	if err := run(); err != nil {
		slog.Error("fatal", "err", err)
		os.Exit(1)
	}
}

func run() error {
	configPath := flag.String("config", envOr("CONFIG_PATH", "config/config.yaml"), "path to config.yaml")
	migrationsDir := flag.String("migrations", envOr("MIGRATIONS_DIR", "migrations"), "directory with .sql migrations")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg, err := config.Load(*configPath)
	if err != nil {
		return err // fail-fast: never start with a half-valid configuration
	}
	if cfg.IsProduction() && secretLooksLikePlaceholder(cfg) {
		return errors.New("refusing to start production with placeholder secrets in config.yaml")
	}
	logger.Info("config loaded",
		"project", cfg.Project,
		"environment", cfg.Environment,
		"listen", cfg.ListenAddr(),
		"public_url", cfg.Server.PublicURL,
	)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := waitPostgres(ctx, cfg, logger)
	if err != nil {
		return err
	}
	defer db.Close()

	migrations, err := loadMigrations(*migrationsDir)
	if err != nil {
		return err
	}
	for name, script := range migrations {
		if err := db.Migrate(ctx, script); err != nil {
			return fmt.Errorf("migration %s: %w", name, err)
		}
		logger.Info("migration applied", "file", name)
	}

	rdb := store.NewRedis(cfg)
	defer func() { _ = rdb.Close() }()
	if err := waitRedis(ctx, func(c context.Context) error { return rdb.Ping(c).Err() }, logger); err != nil {
		return err
	}

	limiter := defense.NewLimiter(rdb, cfg.Defense.Kotlin.RedisStream, logger)
	hub := ws.NewHub(logger)
	jwtIssuer := auth.NewIssuer(cfg.JWT.Secret, cfg.JWTTTL(), cfg.JWT.Issuer)
	api := httpapi.New(cfg, db, rdb, limiter, hub, jwtIssuer, logger)

	srv := &http.Server{
		Addr:              cfg.ListenAddr(),
		Handler:           api.Handler(),
		ReadHeaderTimeout: cfg.ReadHeaderTimeout(),
		IdleTimeout:       cfg.IdleTimeout(),
		MaxHeaderBytes:    1 << 20,
		ErrorLog:          slog.NewLogLogger(logger.Handler(), slog.LevelWarn),
	}

	serveErr := make(chan error, 1)
	go func() {
		logger.Info("http listening", "addr", cfg.ListenAddr())
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
		}
	}()

	// Kernel-space ban mirror for the XDP loader.
	go defense.SyncBanFile(ctx, rdb, envOr("BAN_FILE", "/var/lib/paperplane/bans.txt"),
		5*time.Second, logger)
	go cleanupLoop(ctx, db, logger)

	select {
	case err := <-serveErr:
		return fmt.Errorf("http server: %w", err)
	case <-ctx.Done():
	}

	logger.Info("shutdown requested")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout())
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Warn("graceful shutdown incomplete", "err", err)
	}
	hub.CloseAll()
	logger.Info("bye")
	return nil
}

func waitPostgres(ctx context.Context, cfg *config.Config, log *slog.Logger) (*store.Postgres, error) {
	var lastErr error
	for attempt := 1; attempt <= 20; attempt++ {
		db, err := store.NewPostgres(ctx, cfg.PostgresDSN())
		if err == nil {
			pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
			err = db.Ping(pingCtx)
			cancel()
			if err == nil {
				log.Info("postgres connected", "attempt", attempt)
				return db, nil
			}
			db.Close()
		}
		lastErr = err
		log.Warn("postgres not ready, retrying", "attempt", attempt, "err", err)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
	return nil, fmt.Errorf("postgres unreachable: %w", lastErr)
}

func waitRedis(ctx context.Context, ping func(context.Context) error, log *slog.Logger) error {
	var lastErr error
	for attempt := 1; attempt <= 20; attempt++ {
		pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		err := ping(pingCtx)
		cancel()
		if err == nil {
			log.Info("redis connected", "attempt", attempt)
			return nil
		}
		lastErr = err
		log.Warn("redis not ready, retrying", "attempt", attempt, "err", err)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
	return fmt.Errorf("redis unreachable: %w", lastErr)
}

func cleanupLoop(ctx context.Context, db *store.Postgres, log *slog.Logger) {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			runCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			n, err := db.CleanupExpiredEnvelopes(runCtx)
			cancel()
			if err != nil {
				log.Warn("envelope cleanup failed", "err", err)
				continue
			}
			if n > 0 {
				log.Info("expired envelopes removed", "count", n)
			}
		}
	}
}

func loadMigrations(dir string) (map[string]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read migrations dir %q: %w", dir, err)
	}
	out := map[string]string{}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		// Map iteration is random; migrations here are idempotent and ordered
		// by a single file, so name keying is enough for this project.
		out[e.Name()] = string(raw)
	}
	if len(out) == 0 {
		return nil, errors.New("no .sql migrations found")
	}
	return out, nil
}

func secretLooksLikePlaceholder(cfg *config.Config) bool {
	joined := cfg.JWT.Secret + cfg.Crypto.Pepper + cfg.Database.Password + cfg.Redis.Password
	return strings.Contains(joined, "change-me")
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
