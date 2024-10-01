-- +goose Up
-- +goose StatementBegin

CREATE TABLE configs
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE
);

CREATE TABLE configs_values
(
    config_id INTEGER REFERENCES configs (id),
    key       TEXT DEFAULT '',
    value     TEXT DEFAULT '',
    UNIQUE (config_id, key)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE configs_values;
DROP TABLE configs;
-- +goose StatementEnd
