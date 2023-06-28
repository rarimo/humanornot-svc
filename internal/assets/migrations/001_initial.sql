-- +migrate Up

CREATE TABLE users(
    id            UUID         PRIMARY KEY    NOT NULL,
    status        VARCHAR(256)                NOT NULL,
    created_at    TIMESTAMP    WITH TIME ZONE NOT NULL,
    identity_id   BYTEA        UNIQUE         NOT NULL,
    eth_address   BYTEA        UNIQUE,
    provider_data BYTEA
);

-- +migrate Down

DROP TABLE users;