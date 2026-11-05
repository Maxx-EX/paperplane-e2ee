-- Paperplane schema. No plaintext, no reversible username, no long-term ciphertext.
CREATE TABLE IF NOT EXISTS accounts (
    account_id   TEXT PRIMARY KEY,
    auth_hash    TEXT NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    last_seen_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS prekey_bundles (
    account_id             TEXT PRIMARY KEY,
    identity_key           BYTEA NOT NULL,
    signed_prekey_id       BIGINT NOT NULL,
    signed_prekey_public   BYTEA NOT NULL,
    signed_prekey_signature BYTEA NOT NULL,
    kyber_prekey_id        BIGINT,
    kyber_prekey_public    BYTEA,
    kyber_prekey_signature BYTEA,
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS one_time_prekeys (
    account_id TEXT NOT NULL,
    prekey_id  BIGINT NOT NULL,
    public_key BYTEA NOT NULL,
    PRIMARY KEY (account_id, prekey_id)
);

CREATE TABLE IF NOT EXISTS envelopes (
    id            BIGSERIAL PRIMARY KEY,
    target_id     TEXT NOT NULL,
    sender_id     TEXT,
    cipher_type   SMALLINT NOT NULL,
    body          BYTEA NOT NULL,
    client_msg_id TEXT NOT NULL,
    sent_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    expires_at    TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS envelopes_target_idx ON envelopes (target_id, id);

CREATE INDEX IF NOT EXISTS envelopes_expiry_idx ON envelopes (expires_at);

CREATE INDEX IF NOT EXISTS one_time_prekeys_account_idx ON one_time_prekeys (account_id);
