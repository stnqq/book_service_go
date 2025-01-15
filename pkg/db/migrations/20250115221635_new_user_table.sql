-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    author VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd