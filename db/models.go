// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Post struct {
	ID        pgtype.UUID      `json:"id"`
	Title     pgtype.Text      `json:"title"`
	Content   pgtype.Text      `json:"content"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}
