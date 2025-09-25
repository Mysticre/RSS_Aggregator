-- +goose  UP

ALTER TABLE feeds ADD last_fetch_at TIMESTAMP DEFAULT NOW() NOT NULL;



-- +goose DOWN

ALTER TABLE feeds DROP COLUMN last_fetch_at;