-- +goose Up
-- +goose StatementBegin
CREATE TABLE configs (
    key TEXT,
    value TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE configs;
-- +goose StatementEnd
