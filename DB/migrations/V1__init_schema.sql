CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS auth_users (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS storages (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    owner_user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    path TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_storages_owner_user FOREIGN KEY (owner_user_id) REFERENCES auth_users (id) ON DELETE CASCADE,
    CONSTRAINT uq_storages_owner_name UNIQUE (owner_user_id, name),
    CONSTRAINT uq_storages_path UNIQUE (path)
);

CREATE INDEX IF NOT EXISTS idx_storages_owner_user_id ON storages (owner_user_id);
CREATE INDEX IF NOT EXISTS idx_storages_updated_at ON storages (updated_at);
