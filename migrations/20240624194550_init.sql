-- +goose Up
-- +goose StatementBegin

CREATE TABLE configs
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE configs_values
(
    config_id INTEGER REFERENCES configs (id),
    key       TEXT DEFAULT '',
    value     TEXT DEFAULT '',
    version   TEXT DEFAULT '',
    UNIQUE (config_id, key, version)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE configs_values;
DROP TABLE configs;
-- +goose StatementEnd
