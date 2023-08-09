-- +migrate Up

CREATE TABLE users(
    id            UUID         PRIMARY KEY       NOT NULL,
    status        VARCHAR(256)                   NOT NULL,
    created_at    TIMESTAMP    WITHOUT TIME ZONE NOT NULL,
    identity_id   BYTEA        UNIQUE            NOT NULL,
    eth_address   BYTEA,
    provider_data BYTEA,
    provider_hash BYTEA        UNIQUE
);

CREATE TABLE nonce(
    id          BIGSERIAL PRIMARY KEY       NOT NULL,
    nonce       TEXT                        NOT NULL,
    expires_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    eth_address BYTEA                       NOT NULL
);

-- +migrate Down

DROP TABLE users;
drop table nonce;