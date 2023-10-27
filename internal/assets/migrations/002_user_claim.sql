-- +migrate Up

ALTER TABLE users ADD COLUMN claim_id UUID;

-- +migrate Down

ALTER TABLE users DROP COLUMN claim_id;