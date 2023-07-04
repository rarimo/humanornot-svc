-- +migrate Up
create table nonce
(
    id         BIGSERIAL PRIMARY KEY,
    message    TEXT   NOT NULL,
    expires_at BIGINT NOT NULL,
    address    BYTEA  NOT NULL
);

-- +migrate Down

drop table nonce;
