-- +goose Up
-- +goose StatementBegin
ALTER TABLE configs_values
    ADD COLUMN version TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE configs_values
    DROP COLUMN version;
-- +goose StatementEnd
