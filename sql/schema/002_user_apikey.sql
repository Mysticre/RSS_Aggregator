-- +goose UP

CREATE EXTENSION IF NOT EXISTS pgcrypto; --postgreSQL 引用套件方法


ALTER TABLE users ADD api_key VARCHAR(64) NOT NULL UNIQUE
DEFAULT
encode(digest(random()::text, 'sha256'), 'hex');





-- +goose DOWN


ALTER TABLE users 
DROP api_key;