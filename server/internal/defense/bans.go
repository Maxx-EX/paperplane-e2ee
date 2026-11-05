package defense

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// SyncBanFile mirrors the Redis ban ZSET to a plain text file that the XDP
// loader watches. Kernel-space filtering therefore keeps working even when the
// userspace loader and Redis are briefly out of sync; the file is the only
// contract between Go and C.
//
// Format: one "ip expiry_unix" pair per line.
func SyncBanFile(ctx context.Context, rdb *redis.Client, path string, every time.Duration, log *slog.Logger) {
	if path == "" {
		return
	}
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	var lastErr string
	write := func() {
		now := time.Now().Unix()
		ips, err := rdb.ZRangeByScore(ctx, banKey, &redis.ZRangeBy{
			Min: "(" + strconv.FormatInt(now, 10),
			Max: "+inf",
		}).Result()
		if err != nil {
			log.Warn("ban sync read failed", "err", err)
			return
		}
		var b strings.Builder
		for _, ip := range ips {
			score, err := rdb.ZScore(ctx, banKey, ip).Result()
			if err != nil {
				continue
			}
			fmt.Fprintf(&b, "%s %d\n", ip, int64(score))
		}
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			if lastErr != err.Error() {
				log.Warn("ban file directory unavailable", "path", path, "err", err)
				lastErr = err.Error()
			}
			return
		}
		tmp := path + ".tmp"
		if err := os.WriteFile(tmp, []byte(b.String()), 0o644); err != nil {
			if lastErr != err.Error() {
				log.Warn("ban file write failed", "err", err)
				lastErr = err.Error()
			}
			return
		}
		if err := os.Rename(tmp, path); err != nil {
			log.Warn("ban file rename failed", "err", err)
			return
		}
		lastErr = ""
	}

	write()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			write()
		}
	}
}
