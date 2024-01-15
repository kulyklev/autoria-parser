-- +goose Up
CREATE TABLE failed_parses
(
    id         UUID PRIMARY KEY,
    auto_id    INT       NOT NULL,
    err        VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS failed_parses;
