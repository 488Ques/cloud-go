-- +goose Up
-- +goose StatementBegin
CREATE TABLE Post (
    id UUID PRIMARY KEY,
    title VARCHAR(100),
    content TEXT,
    created_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Post;
-- +goose StatementEnd
