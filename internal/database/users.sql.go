// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one
    INSERT INTO users (id, created_at , updated_at, name)
    VALUES (s1,s2 ,s3 ,s4)
    RETURNING id, created_at, updated_at, name
`

func (q *Queries) CreateUser(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}