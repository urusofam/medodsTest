CREATE DATABASE authDB;

CREATE TABLE IF NOT EXISTS users (
    guid UUID PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(guid) ON DELETE CASCADE,
    token_hash TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    ip INET NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user_id ON refresh_tokens(user_id);