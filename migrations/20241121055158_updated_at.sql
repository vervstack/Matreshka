-- +goose Up
-- +goose StatementBegin
ALTER TABLE configs
    ADD COLUMN updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP;
UPDATE configs SET updated_at = current_timestamp;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE configs DROP COLUMN updated_at;
-- +goose StatementEnd
