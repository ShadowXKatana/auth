CREATE TABLE IF NOT EXISTS auth_refresh_tokens (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL,
    token_hash TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_auth_refresh_tokens_user FOREIGN KEY (user_id) REFERENCES auth_users (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_auth_refresh_tokens_user_id ON auth_refresh_tokens (user_id);
CREATE INDEX IF NOT EXISTS idx_auth_refresh_tokens_expires_at ON auth_refresh_tokens (expires_at);

CREATE TABLE IF NOT EXISTS auth_user_profiles (
    user_id UUID PRIMARY KEY,
    display_name VARCHAR(255),
    avatar_url TEXT,
    bio TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_auth_user_profiles_user FOREIGN KEY (user_id) REFERENCES auth_users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS storage_items (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    storage_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    value NUMERIC(18,2),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_storage_items_storage FOREIGN KEY (storage_id) REFERENCES storages (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_storage_items_storage_id ON storage_items (storage_id);
CREATE INDEX IF NOT EXISTS idx_storage_items_name ON storage_items (name);

CREATE TABLE IF NOT EXISTS storage_tags (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS storage_item_tags (
    storage_item_id UUID NOT NULL,
    storage_tag_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT pk_storage_item_tags PRIMARY KEY (storage_item_id, storage_tag_id),
    CONSTRAINT fk_storage_item_tags_item FOREIGN KEY (storage_item_id) REFERENCES storage_items (id) ON DELETE CASCADE,
    CONSTRAINT fk_storage_item_tags_tag FOREIGN KEY (storage_tag_id) REFERENCES storage_tags (id) ON DELETE CASCADE
);
