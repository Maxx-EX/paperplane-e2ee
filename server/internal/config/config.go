// Package config loads and validates the single source of truth:
// config/config.yaml. Every service reads the same file; nothing else is
// allowed to carry its own copy of server settings.
package config

import (
	"bytes"
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Config mirrors config/config.yaml field for field.
type Config struct {
	Project     string   `yaml:"project"`
	Environment string   `yaml:"environment"`
	Server      Server   `yaml:"server"`
	Database    Database `yaml:"database"`
	Redis       Redis    `yaml:"redis"`
	JWT         JWT      `yaml:"jwt"`
	Crypto      Crypto   `yaml:"crypto"`
	Client      Client   `yaml:"client"`
	Defense     Defense  `yaml:"defense"`
}

type Server struct {
	Host              string   `yaml:"host"`
	Port              int      `yaml:"port"`
	PublicURL         string   `yaml:"public_url"`
	WSAllowedOrigins  []string `yaml:"ws_allowed_origins"`
	ReadHeaderTimeout string   `yaml:"read_header_timeout"`
	IdleTimeout       string   `yaml:"idle_timeout"`
	ShutdownTimeout   string   `yaml:"shutdown_timeout"`
}

type Database struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Name            string `yaml:"name"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	MaxConns        int32  `yaml:"max_conns"`
	EnvelopeTTLHour int    `yaml:"envelope_ttl_hours"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
}

type JWT struct {
	Secret string `yaml:"secret"`
	TTL    string `yaml:"ttl"`
	Issuer string `yaml:"issuer"`
}

type Crypto struct {
	Pepper string `yaml:"pepper"`
	Argon2 struct {
		MemoryKiB   uint32 `yaml:"memory_kib"`
		Iterations  uint32 `yaml:"iterations"`
		Parallelism uint8  `yaml:"parallelism"`
	} `yaml:"argon2"`
}

type Client struct {
	DevServerURL string `yaml:"dev_server_url"`
	AppName      string `yaml:"app_name"`
}

type Defense struct {
	XDP        XDP        `yaml:"xdp"`
	RateLimits RateLimits `yaml:"rate_limits"`
	Kotlin     Kotlin     `yaml:"kotlin"`
	CrowdSec   CrowdSec   `yaml:"crowdsec"`
}

type XDP struct {
	Enabled bool   `yaml:"enabled"`
	Iface   string `yaml:"iface"`
}

type RateLimits struct {
	WSPerIP                 int    `yaml:"ws_per_ip"`
	PreKeyPerMin            int    `yaml:"prekey_per_min"`
	EnvelopePerMin          int    `yaml:"envelope_per_min"`
	EnvelopeDistinctTargets int    `yaml:"envelope_distinct_targets"`
	EnvelopeWindow          string `yaml:"envelope_window"`
	ChallengePerMin         int    `yaml:"challenge_per_min"`
}

type Kotlin struct {
	RedisStream   string `yaml:"redis_stream"`
	ConsumerGroup string `yaml:"consumer_group"`
	Window        string `yaml:"window"`
}

type CrowdSec struct {
	Collections []string `yaml:"collections"`
	LAPIPort    int      `yaml:"lapi_port"`
}

// Load reads config.yaml, rejects unknown keys and fails fast on invalid input.
func Load(path string) (*Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", path, err)
	}
	var c Config
	dec := yaml.NewDecoder(bytes.NewReader(raw))
	dec.KnownFields(true) // typo in config.yaml is a startup error, not a silent default
	if err := dec.Decode(&c); err != nil {
		return nil, fmt.Errorf("parse config %q: %w", path, err)
	}
	c.applyDefaults()
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Config) applyDefaults() {
	if c.Server.ReadHeaderTimeout == "" {
		c.Server.ReadHeaderTimeout = "10s"
	}
	if c.Server.IdleTimeout == "" {
		c.Server.IdleTimeout = "120s"
	}
	if c.Server.ShutdownTimeout == "" {
		c.Server.ShutdownTimeout = "15s"
	}
	if c.JWT.TTL == "" {
		c.JWT.TTL = "24h"
	}
	if c.JWT.Issuer == "" {
		c.JWT.Issuer = c.Project
	}
	if c.Database.MaxConns == 0 {
		c.Database.MaxConns = 16
	}
	if c.Database.EnvelopeTTLHour == 0 {
		c.Database.EnvelopeTTLHour = 24
	}
	if c.Defense.RateLimits.EnvelopeWindow == "" {
		c.Defense.RateLimits.EnvelopeWindow = "60s"
	}
	if c.Defense.Kotlin.RedisStream == "" {
		c.Defense.Kotlin.RedisStream = "paperplane:events"
	}
	if c.Defense.Kotlin.ConsumerGroup == "" {
		c.Defense.Kotlin.ConsumerGroup = "paperplane-policy"
	}
	if c.Defense.Kotlin.Window == "" {
		c.Defense.Kotlin.Window = "60s"
	}
	if c.Defense.CrowdSec.LAPIPort == 0 {
		c.Defense.CrowdSec.LAPIPort = 8081
	}
}

// Validate enforces every field the runtime depends on. Missing or nonsensical
// values abort startup instead of surfacing as a confusing runtime bug.
func (c *Config) Validate() error {
	problems := make([]string, 0, 8)

	if strings.TrimSpace(c.Project) == "" {
		problems = append(problems, "project is required")
	}
	switch c.Environment {
	case "production", "development":
	default:
		problems = append(problems, `environment must be "production" or "development"`)
	}
	if c.Server.Host == "" {
		problems = append(problems, "server.host is required")
	}
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		problems = append(problems, "server.port must be 1-65535")
	}
	u, err := url.Parse(c.Server.PublicURL)
	if err != nil || u.Host == "" {
		problems = append(problems, "server.public_url must be an absolute URL")
	} else if c.Environment == "production" && u.Scheme != "https" {
		problems = append(problems, "server.public_url must use https:// in production")
	}
	if len(c.Server.WSAllowedOrigins) == 0 {
		problems = append(problems, "server.ws_allowed_origins must not be empty")
	}
	for _, o := range c.Server.WSAllowedOrigins {
		if _, err := url.Parse(o); err != nil || !strings.Contains(o, "://") {
			problems = append(problems, "invalid ws origin: "+o)
		}
	}
	if c.Database.Host == "" || c.Database.Name == "" || c.Database.User == "" {
		problems = append(problems, "database host/name/user are required")
	}
	if c.Database.Port < 1 || c.Database.Port > 65535 {
		problems = append(problems, "database.port must be 1-65535")
	}
	if c.Database.Password == "" {
		problems = append(problems, "database.password is required")
	}
	if c.Redis.Host == "" {
		problems = append(problems, "redis.host is required")
	}
	if c.Redis.Port < 1 || c.Redis.Port > 65535 {
		problems = append(problems, "redis.port must be 1-65535")
	}
	if len(c.JWT.Secret) < 32 {
		problems = append(problems, "jwt.secret must be at least 32 bytes")
	}
	if len(c.Crypto.Pepper) < 16 {
		problems = append(problems, "crypto.pepper must be at least 16 bytes")
	}
	if c.Crypto.Argon2.MemoryKiB < 19456 {
		problems = append(problems, "crypto.argon2.memory_kib must be >= 19456 (19 MiB)")
	}
	if c.Crypto.Argon2.Iterations < 1 || c.Crypto.Argon2.Iterations > 10 {
		problems = append(problems, "crypto.argon2.iterations must be 1-10")
	}
	if c.Crypto.Argon2.Parallelism < 1 {
		problems = append(problems, "crypto.argon2.parallelism must be >= 1")
	}
	if c.Defense.RateLimits.WSPerIP < 1 || c.Defense.RateLimits.PreKeyPerMin < 1 ||
		c.Defense.RateLimits.EnvelopePerMin < 1 || c.Defense.RateLimits.ChallengePerMin < 1 {
		problems = append(problems, "defense.rate_limits values must be positive")
	}
	if c.Defense.XDP.Enabled && c.Defense.XDP.Iface == "" {
		problems = append(problems, "defense.xdp.iface is required when xdp is enabled")
	}
	if _, err := time.ParseDuration(c.JWT.TTL); err != nil {
		problems = append(problems, "jwt.ttl is not a duration (e.g. 24h)")
	}
	for _, d := range []struct{ name, value string }{
		{"server.read_header_timeout", c.Server.ReadHeaderTimeout},
		{"server.idle_timeout", c.Server.IdleTimeout},
		{"server.shutdown_timeout", c.Server.ShutdownTimeout},
		{"defense.rate_limits.envelope_window", c.Defense.RateLimits.EnvelopeWindow},
		{"defense.kotlin.window", c.Defense.Kotlin.Window},
	} {
		if _, err := time.ParseDuration(d.value); err != nil {
			problems = append(problems, d.name+" is not a duration")
		}
	}
	if len(problems) > 0 {
		return fmt.Errorf("invalid config: %s", strings.Join(problems, "; "))
	}
	return nil
}

func (c *Config) IsProduction() bool { return c.Environment == "production" }

func (c *Config) ListenAddr() string {
	return net.JoinHostPort(c.Server.Host, strconv.Itoa(c.Server.Port))
}

func (c *Config) PostgresDSN() string {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.Database.User, c.Database.Password),
		Host:   net.JoinHostPort(c.Database.Host, strconv.Itoa(c.Database.Port)),
		Path:   "/" + c.Database.Name,
	}
	q := u.Query()
	q.Set("sslmode", "disable")
	q.Set("pool_max_conns", strconv.Itoa(int(c.Database.MaxConns)))
	q.Set("application_name", c.Project+"-server")
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *Config) RedisAddr() string {
	return net.JoinHostPort(c.Redis.Host, strconv.Itoa(c.Redis.Port))
}

func (c *Config) JWTTTL() time.Duration { return mustDuration(c.JWT.TTL, 24*time.Hour) }
func (c *Config) ReadHeaderTimeout() time.Duration {
	return mustDuration(c.Server.ReadHeaderTimeout, 10*time.Second)
}
func (c *Config) IdleTimeout() time.Duration {
	return mustDuration(c.Server.IdleTimeout, 2*time.Minute)
}
func (c *Config) ShutdownTimeout() time.Duration {
	return mustDuration(c.Server.ShutdownTimeout, 15*time.Second)
}
func (c *Config) EnvelopeWindow() time.Duration {
	return mustDuration(c.Defense.RateLimits.EnvelopeWindow, time.Minute)
}
func (c *Config) PolicyWindow() time.Duration {
	return mustDuration(c.Defense.Kotlin.Window, time.Minute)
}

func mustDuration(s string, fallback time.Duration) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return fallback
	}
	return d
}
