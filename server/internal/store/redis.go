package store

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/paperplane/paperplane/server/internal/config"
)

// NewRedis builds the shared Redis client. Redis holds only ephemeral state:
// nonces, WebSocket tickets, rate-limit counters, ban decisions and the
// defence event stream. No ciphertext and no credential is ever written here.
func NewRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         cfg.RedisAddr(),
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: 4,
		DialTimeout:  3,
	})
}

func PingRedis(ctx context.Context, rdb *redis.Client) error {
	if err := rdb.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis ping: %w", err)
	}
	return nil
}
