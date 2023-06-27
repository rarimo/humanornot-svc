-- +migrate Up

CREATE TABLE users(
    id            UUID         PRIMARY KEY  NOT NULL,
    status        VARCHAR(256)              NOT NULL,
    identity_id   BYTEA                     NOT NULL,
    eth_address   BYTEA        UNIQUE,
    provider_data BYTEA
);

-- +migrate Down

DROP TABLE users;