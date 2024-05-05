-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY title;

-- name: CreateBook :one
INSERT INTO books (
    id,
    title,
    author,
    published_date,
    image_url,
    description
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateBooks :one
UPDATE books
  SET title = $2,
  author = $3,
  published_date = $4,
  image_url = $5,
  description = $6
WHERE id = $1
RETURNING *;

-- name: DeleteBooks :exec
DELETE FROM books
WHERE id = $1;
