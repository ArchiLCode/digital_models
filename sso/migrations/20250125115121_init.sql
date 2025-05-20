-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id UUID PRIMARY KEY,
    name     VARCHAR(60) NOT NULL,
    last_name     VARCHAR(60) NOT NULL,
    goal     VARCHAR(60) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    pass_hash BYTEA NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_email ON users (email);

CREATE TABLE IF NOT EXISTS apps
(
    id     INTEGER PRIMARY KEY,
    name   TEXT NOT NULL UNIQUE,
    secret TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS admins
(
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);

-- Chats table
CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    is_private BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Chat participants table
CREATE TABLE IF NOT EXISTS chat_participants (
    chat_id INTEGER REFERENCES chats(id),
    user_id UUID REFERENCES users(id),  
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (chat_id, user_id)
);

-- Messages table
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    chat_id INTEGER REFERENCES chats(id),
    user_id UUID REFERENCES users(id), 
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP DATABASE models_db;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS apps CASCADE;
DROP TABLE IF EXISTS admins CASCADE;
DROP TABLE IF EXISTS goose_db_version;

-- +goose StatementEnd
