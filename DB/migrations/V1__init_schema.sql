CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE "user" (
    id   UUID         NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE storage (
    id          UUID         NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id     UUID         NOT NULL UNIQUE,
    name        VARCHAR(255) NOT NULL,
    last_update TIMESTAMP,
    CONSTRAINT fk_storage_user FOREIGN KEY (user_id) REFERENCES "user" (id)
);

CREATE TABLE item (
    id         UUID         NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    storage_id UUID         NOT NULL,
    name       VARCHAR(255) NOT NULL,
    value      DECIMAL,
    CONSTRAINT fk_item_storage FOREIGN KEY (storage_id) REFERENCES storage (id)
);

CREATE INDEX idx_item_name ON item (name);

CREATE TABLE tag (
    id   UUID         NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) UNIQUE
);

CREATE TABLE item_tag (
    item_id UUID NOT NULL,
    tag_id  UUID NOT NULL,
    CONSTRAINT pk_item_tag      PRIMARY KEY (item_id, tag_id),
    CONSTRAINT fk_item_tag_item FOREIGN KEY (item_id) REFERENCES item (id),
    CONSTRAINT fk_item_tag_tag  FOREIGN KEY (tag_id)  REFERENCES tag (id)
);
