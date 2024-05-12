-- name: GetPost :one
SELECT * FROM Post
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM Post
ORDER BY title;

-- name: CreatePost :one
INSERT INTO Post (
    id,
    title,
    content
) VALUES (
    $1, $2, $3
)
RETURNING id, title, content;

-- name: UpdatePosts :one
UPDATE Post
  SET title = $2,
  content = $3
WHERE id = $1
RETURNING *;

-- name: DeletePosts :exec
DELETE FROM Post
WHERE id = $1;
