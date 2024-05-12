-- +goose Up
-- +goose StatementBegin
CREATE TABLE Post (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Post;
-- +goose StatementEnd
