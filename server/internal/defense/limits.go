// Package defense implements the embedded, in-process defence layer:
// token buckets, distinct-target sliding windows, ban lookups and the bridge
// that feeds kernel-space bans to the XDP loader.
//
// Every Redis call in the hot path is bounded and fails open. A Redis outage
// must degrade protection, never take the chat service down.
package defense

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const banKey = "paperplane:ban:ip"

type Limiter struct {
	rdb    *redis.Client
	log    *slog.Logger
	stream string
}

func NewLimiter(rdb *redis.Client, stream string, log *slog.Logger) *Limiter {
	return &Limiter{rdb: rdb, log: log, stream: stream}
}

var tokenBucketScript = redis.NewScript(`
local key = KEYS[1]
local rate = tonumber(ARGV[1])
local burst = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local pair = redis.call('HMGET', key, 'tokens', 'ts')
local tokens = tonumber(pair[1])
local ts = tonumber(pair[2])
if tokens == nil then tokens = burst; ts = now end
local elapsed = now - ts
if elapsed < 0 then elapsed = 0 end
tokens = math.min(burst, tokens + elapsed * rate / 60)
local allowed = 0
if tokens >= 1 then
  tokens = tokens - 1
  allowed = 1
end
redis.call('HSET', key, 'tokens', tokens, 'ts', now)
redis.call('EXPIRE', key, math.ceil(burst * 60 / rate) + 60)
return allowed
`)

// AllowTokenBucket consumes one token for subject in the given bucket.
// Fail-open on Redis errors (logged), so an outage cannot lock users out.
func (l *Limiter) AllowTokenBucket(ctx context.Context, bucket, subject string, perMinute int) bool {
	if perMinute <= 0 {
		return true
	}
	key := "paperplane:rl:" + bucket + ":" + subject
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	n, err := tokenBucketScript.Run(ctx, l.rdb, []string{key},
		perMinute, perMinute, time.Now().Unix()).Int()
	if err != nil {
		l.log.Warn("rate limiter unavailable, failing open", "bucket", bucket, "err", err)
		return true
	}
	return n == 1
}

var distinctTargetsScript = redis.NewScript(`
local key = KEYS[1]
local now = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local limit = tonumber(ARGV[3])
local member = ARGV[4]
redis.call('ZREMRANGEBYSCORE', key, 0, now - window)
redis.call('ZADD', key, now, member)
local count = redis.call('ZCARD', key)
redis.call('EXPIRE', key, window + 5)
if count > limit then return 0 end
return 1
`)

// AllowDistinctTargets throttles fan-out: a compromised account blasting many
// recipients inside one window is the classic spam / harassment signature.
func (l *Limiter) AllowDistinctTargets(ctx context.Context, subject, target string, limit int, window time.Duration) bool {
	if limit <= 0 {
		return true
	}
	key := "paperplane:fanout:" + subject
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	n, err := distinctTargetsScript.Run(ctx, l.rdb, []string{key},
		time.Now().Unix(), int(window.Seconds()), limit, target).Int()
	if err != nil {
		l.log.Warn("fan-out limiter unavailable, failing open", "err", err)
		return true
	}
	return n == 1
}

// IsBanned reads the ban ZSET written by the Kotlin policy sidecar.
func (l *Limiter) IsBanned(ctx context.Context, ip string) bool {
	if ip == "" {
		return false
	}
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	score, err := l.rdb.ZScore(ctx, banKey, ip).Result()
	if errors.Is(err, redis.Nil) {
		return false
	}
	if err != nil {
		l.log.Warn("ban lookup failed, failing open", "err", err)
		return false
	}
	return score > float64(time.Now().Unix())
}

// PublishEvent feeds the Kotlin sidecar through a capped Redis Stream.
// Only identifiers and coarse counters are published — never ciphertext.
func (l *Limiter) PublishEvent(ctx context.Context, kind string, fields map[string]string) {
	if l.stream == "" {
		return
	}
	values := map[string]any{
		"kind": kind,
		"ts":   strconv.FormatInt(time.Now().Unix(), 10),
	}
	for k, v := range fields {
		values[k] = v
	}
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	err := l.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: l.stream,
		MaxLen: 20000,
		Approx: true,
		Values: values,
	}).Err()
	if err != nil {
		l.log.Warn("event stream unavailable", "kind", kind, "err", err)
	}
}

// CountReconnect increments a short-lived counter used by the sidecar's
// reconnect-velocity rule.
func (l *Limiter) CountReconnect(ctx context.Context, accountID string) int64 {
	key := "paperplane:reconnect:" + accountID
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	n, err := l.rdb.Incr(ctx, key).Result()
	if err != nil {
		return 0
	}
	_ = l.rdb.Expire(ctx, key, 30*time.Second).Err()
	return n
}
