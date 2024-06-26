// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBook = `-- name: CreateBook :one
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
RETURNING id, title, author, published_date, image_url, description, created_at, updated_at, deleted_at
`

type CreateBookParams struct {
	ID            pgtype.UUID `json:"id"`
	Title         string      `json:"title"`
	Author        string      `json:"author"`
	PublishedDate pgtype.Date `json:"published_date"`
	ImageUrl      pgtype.Text `json:"image_url"`
	Description   pgtype.Text `json:"description"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRow(ctx, createBook,
		arg.ID,
		arg.Title,
		arg.Author,
		arg.PublishedDate,
		arg.ImageUrl,
		arg.Description,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.PublishedDate,
		&i.ImageUrl,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteBooks = `-- name: DeleteBooks :exec
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBooks(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteBooks, id)
	return err
}

const getBook = `-- name: GetBook :one
SELECT id, title, author, published_date, image_url, description, created_at, updated_at, deleted_at FROM books
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBook(ctx context.Context, id pgtype.UUID) (Book, error) {
	row := q.db.QueryRow(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.PublishedDate,
		&i.ImageUrl,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many
SELECT id, title, author, published_date, image_url, description, created_at, updated_at, deleted_at FROM books
ORDER BY title
`

func (q *Queries) ListBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.Query(ctx, listBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.PublishedDate,
			&i.ImageUrl,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBooks = `-- name: UpdateBooks :one
UPDATE books
  SET title = $2,
  author = $3,
  published_date = $4,
  image_url = $5,
  description = $6
WHERE id = $1
RETURNING id, title, author, published_date, image_url, description, created_at, updated_at, deleted_at
`

type UpdateBooksParams struct {
	ID            pgtype.UUID `json:"id"`
	Title         string      `json:"title"`
	Author        string      `json:"author"`
	PublishedDate pgtype.Date `json:"published_date"`
	ImageUrl      pgtype.Text `json:"image_url"`
	Description   pgtype.Text `json:"description"`
}

func (q *Queries) UpdateBooks(ctx context.Context, arg UpdateBooksParams) (Book, error) {
	row := q.db.QueryRow(ctx, updateBooks,
		arg.ID,
		arg.Title,
		arg.Author,
		arg.PublishedDate,
		arg.ImageUrl,
		arg.Description,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.PublishedDate,
		&i.ImageUrl,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
