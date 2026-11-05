// Package store contains the durable state layer. PostgreSQL holds accounts,
// pre-key bundles and undelivered ciphertext envelopes; nothing here ever
// stores a plaintext message, a password or a reversible username.
package store

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAccountExists = errors.New("account already exists")
)

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(ctx context.Context, dsn string) (*Postgres, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse postgres dsn: %w", err)
	}
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.HealthCheckPeriod = 30 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}
	return &Postgres{pool: pool}, nil
}

func (p *Postgres) Ping(ctx context.Context) error { return p.pool.Ping(ctx) }
func (p *Postgres) Close()                         { p.pool.Close() }

// Migrate applies idempotent DDL. Statements are split on ';' because the
// extended protocol used by pgx accepts only one statement per Exec.
func (p *Postgres) Migrate(ctx context.Context, script string) error {
	for _, stmt := range strings.Split(script, ";") {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" || strings.HasPrefix(stmt, "--") && !strings.Contains(stmt, "\n") {
			continue
		}
		if _, err := p.pool.Exec(ctx, stmt); err != nil {
			return fmt.Errorf("migrate statement %.60q: %w", stmt, err)
		}
	}
	return nil
}

// ---------------------------------------------------------------- accounts

func (p *Postgres) CreateAccount(ctx context.Context, accountID, authHash string) error {
	tag, err := p.pool.Exec(ctx,
		`INSERT INTO accounts (account_id, auth_hash) VALUES ($1, $2)
		 ON CONFLICT (account_id) DO NOTHING`, accountID, authHash)
	if err != nil {
		return fmt.Errorf("insert account: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrAccountExists
	}
	return nil
}

func (p *Postgres) AuthHash(ctx context.Context, accountID string) (string, error) {
	var hash string
	err := p.pool.QueryRow(ctx, `SELECT auth_hash FROM accounts WHERE account_id = $1`, accountID).Scan(&hash)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("select auth hash: %w", err)
	}
	return hash, nil
}

func (p *Postgres) AccountExists(ctx context.Context, accountID string) (bool, error) {
	var exists bool
	err := p.pool.QueryRow(ctx,
		`SELECT EXISTS (SELECT 1 FROM accounts WHERE account_id = $1)`, accountID).Scan(&exists)
	return exists, err
}

func (p *Postgres) TouchAccount(ctx context.Context, accountID string) {
	_, _ = p.pool.Exec(ctx, `UPDATE accounts SET last_seen_at = now() WHERE account_id = $1`, accountID)
}

// ---------------------------------------------------------------- prekeys

type OneTimePreKey struct {
	ID        int64
	PublicKey []byte
}

// Bundle is the public half of a Signal identity. Kyber fields are optional
// because pure-Dart libsignal implements X3DH; the wire format reserves them so
// a PQXDH-capable client can populate them without a schema migration.
type Bundle struct {
	AccountID             string
	IdentityKey           []byte
	SignedPreKeyID        int64
	SignedPreKeyPublic    []byte
	SignedPreKeySignature []byte
	KyberPreKeyID         int64
	KyberPreKeyPublic     []byte
	KyberPreKeySignature  []byte
	OneTimePreKey         *OneTimePreKey
}

func (p *Postgres) UpsertBundle(ctx context.Context, b Bundle, oneTime []OneTimePreKey) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	_, err = tx.Exec(ctx, `
		INSERT INTO prekey_bundles (
			account_id, identity_key, signed_prekey_id, signed_prekey_public,
			signed_prekey_signature, kyber_prekey_id, kyber_prekey_public,
			kyber_prekey_signature, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8, now())
		ON CONFLICT (account_id) DO UPDATE SET
			identity_key = EXCLUDED.identity_key,
			signed_prekey_id = EXCLUDED.signed_prekey_id,
			signed_prekey_public = EXCLUDED.signed_prekey_public,
			signed_prekey_signature = EXCLUDED.signed_prekey_signature,
			kyber_prekey_id = EXCLUDED.kyber_prekey_id,
			kyber_prekey_public = EXCLUDED.kyber_prekey_public,
			kyber_prekey_signature = EXCLUDED.kyber_prekey_signature,
			updated_at = now()`,
		b.AccountID, b.IdentityKey, b.SignedPreKeyID, b.SignedPreKeyPublic,
		b.SignedPreKeySignature, b.KyberPreKeyID, b.KyberPreKeyPublic,
		b.KyberPreKeySignature)
	if err != nil {
		return fmt.Errorf("upsert bundle: %w", err)
	}

	for _, k := range oneTime {
		if _, err := tx.Exec(ctx,
			`INSERT INTO one_time_prekeys (account_id, prekey_id, public_key)
			 VALUES ($1,$2,$3) ON CONFLICT (account_id, prekey_id) DO NOTHING`,
			b.AccountID, k.ID, k.PublicKey); err != nil {
			return fmt.Errorf("insert one-time prekey: %w", err)
		}
	}
	return tx.Commit(ctx)
}

// TakeBundle returns the published bundle and atomically consumes exactly one
// one-time prekey (SKIP LOCKED keeps concurrent fetchers from colliding).
func (p *Postgres) TakeBundle(ctx context.Context, accountID string) (*Bundle, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	b := &Bundle{AccountID: accountID}
	err = tx.QueryRow(ctx, `
		SELECT identity_key,
		       signed_prekey_id, signed_prekey_public, signed_prekey_signature,
		       COALESCE(kyber_prekey_id, 0),
		       COALESCE(kyber_prekey_public, '\x'::bytea),
		       COALESCE(kyber_prekey_signature, '\x'::bytea)
		FROM prekey_bundles WHERE account_id = $1`, accountID).
		Scan(&b.IdentityKey, &b.SignedPreKeyID, &b.SignedPreKeyPublic,
			&b.SignedPreKeySignature, &b.KyberPreKeyID, &b.KyberPreKeyPublic,
			&b.KyberPreKeySignature)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("load bundle: %w", err)
	}

	var otpID int64
	var otpPublic []byte
	err = tx.QueryRow(ctx, `
		SELECT prekey_id, public_key FROM one_time_prekeys
		WHERE account_id = $1 ORDER BY prekey_id LIMIT 1 FOR UPDATE SKIP LOCKED`, accountID).
		Scan(&otpID, &otpPublic)
	switch {
	case err == nil:
		if _, delErr := tx.Exec(ctx,
			`DELETE FROM one_time_prekeys WHERE account_id = $1 AND prekey_id = $2`,
			accountID, otpID); delErr != nil {
			return nil, fmt.Errorf("consume one-time prekey: %w", delErr)
		}
		b.OneTimePreKey = &OneTimePreKey{ID: otpID, PublicKey: otpPublic}
	case errors.Is(err, pgx.ErrNoRows):
		// No OTP left is a normal state: X3DH can proceed without it.
	default:
		return nil, fmt.Errorf("take one-time prekey: %w", err)
	}
	return b, tx.Commit(ctx)
}

func (p *Postgres) CountOneTimePreKeys(ctx context.Context, accountID string) (int, error) {
	var n int
	err := p.pool.QueryRow(ctx,
		`SELECT count(*) FROM one_time_prekeys WHERE account_id = $1`, accountID).Scan(&n)
	return n, err
}

// ---------------------------------------------------------------- envelopes

type Envelope struct {
	ID          int64
	TargetID    string
	SenderID    string // empty when sealed sender
	CipherType  int
	Body        []byte
	ClientMsgID string
	SentAt      time.Time
}

func (p *Postgres) InsertEnvelope(ctx context.Context, e Envelope, ttlHours int) (int64, error) {
	var id int64
	var sender any
	if e.SenderID != "" {
		sender = e.SenderID
	}
	err := p.pool.QueryRow(ctx, `
		INSERT INTO envelopes (target_id, sender_id, cipher_type, body, client_msg_id, expires_at)
		VALUES ($1, $2, $3, $4, $5, now() + make_interval(hours => $6))
		RETURNING id`,
		e.TargetID, sender, e.CipherType, e.Body, e.ClientMsgID, ttlHours).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert envelope: %w", err)
	}
	return id, nil
}

func (p *Postgres) PendingEnvelopes(ctx context.Context, targetID string, limit int) ([]Envelope, error) {
	rows, err := p.pool.Query(ctx, `
		SELECT id, COALESCE(sender_id, ''), cipher_type, body, client_msg_id, sent_at
		FROM envelopes
		WHERE target_id = $1 AND expires_at > now()
		ORDER BY id LIMIT $2`, targetID, limit)
	if err != nil {
		return nil, fmt.Errorf("list envelopes: %w", err)
	}
	defer rows.Close()

	out := make([]Envelope, 0, limit)
	for rows.Next() {
		var e Envelope
		e.TargetID = targetID
		if err := rows.Scan(&e.ID, &e.SenderID, &e.CipherType, &e.Body, &e.ClientMsgID, &e.SentAt); err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, rows.Err()
}

// StripSender drops routing metadata as soon as an envelope reaches a live
// socket. The ciphertext stays only until the recipient acknowledges it.
func (p *Postgres) StripSender(ctx context.Context, ids []int64) {
	if len(ids) == 0 {
		return
	}
	_, _ = p.pool.Exec(ctx, `UPDATE envelopes SET sender_id = NULL WHERE id = ANY($1::bigint[])`, ids)
}

func (p *Postgres) AckEnvelopes(ctx context.Context, targetID string, ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	tag, err := p.pool.Exec(ctx,
		`DELETE FROM envelopes WHERE target_id = $1 AND id = ANY($2::bigint[])`, targetID, ids)
	if err != nil {
		return 0, fmt.Errorf("ack envelopes: %w", err)
	}
	return tag.RowsAffected(), nil
}

func (p *Postgres) CleanupExpiredEnvelopes(ctx context.Context) (int64, error) {
	tag, err := p.pool.Exec(ctx, `DELETE FROM envelopes WHERE expires_at < now()`)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}
