// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: pending.sql

package db

import (
	"context"
	"database/sql"
)

const createPending = `-- name: CreatePending :one
INSERT INTO pending_tasks (
  id,
  title,
  description,
  date,
  is_done,
  is_delete,
  is_favorite
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, title, description, date, is_done, is_delete, is_favorite
`

type CreatePendingParams struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	IsDone      bool   `json:"is_done"`
	IsDelete    bool   `json:"is_delete"`
	IsFavorite  bool   `json:"is_favorite"`
}

func (q *Queries) CreatePending(ctx context.Context, arg CreatePendingParams) (PendingTask, error) {
	row := q.db.QueryRowContext(ctx, createPending,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Date,
		arg.IsDone,
		arg.IsDelete,
		arg.IsFavorite,
	)
	var i PendingTask
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Date,
		&i.IsDone,
		&i.IsDelete,
		&i.IsFavorite,
	)
	return i, err
}

const deletePending = `-- name: DeletePending :exec
DELETE FROM pending_tasks WHERE id = $1
`

func (q *Queries) DeletePending(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePending, id)
	return err
}

const getAllPending = `-- name: GetAllPending :many
SELECT id, title, description, date, is_done, is_delete, is_favorite FROM pending_tasks
ORDER BY date
`

func (q *Queries) GetAllPending(ctx context.Context) ([]PendingTask, error) {
	rows, err := q.db.QueryContext(ctx, getAllPending)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PendingTask{}
	for rows.Next() {
		var i PendingTask
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Date,
			&i.IsDone,
			&i.IsDelete,
			&i.IsFavorite,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPending = `-- name: GetPending :one
SELECT id, title, description, date, is_done, is_delete, is_favorite FROM pending_tasks
WHERE id = $1
`

func (q *Queries) GetPending(ctx context.Context, id string) (PendingTask, error) {
	row := q.db.QueryRowContext(ctx, getPending, id)
	var i PendingTask
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Date,
		&i.IsDone,
		&i.IsDelete,
		&i.IsFavorite,
	)
	return i, err
}

const updatePending = `-- name: UpdatePending :one
UPDATE pending_tasks
SET
   title = COALESCE($4,title),
  description = COALESCE($5,description),
  date = COALESCE($6,date),
  is_done = $1,
  is_delete = $2,
  is_favorite = $3
WHERE
   id = $7
RETURNING id, title, description, date, is_done, is_delete, is_favorite
`

type UpdatePendingParams struct {
	IsDone      bool           `json:"is_done"`
	IsDelete    bool           `json:"is_delete"`
	IsFavorite  bool           `json:"is_favorite"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	Date        sql.NullString `json:"date"`
	ID          string         `json:"id"`
}

func (q *Queries) UpdatePending(ctx context.Context, arg UpdatePendingParams) (PendingTask, error) {
	row := q.db.QueryRowContext(ctx, updatePending,
		arg.IsDone,
		arg.IsDelete,
		arg.IsFavorite,
		arg.Title,
		arg.Description,
		arg.Date,
		arg.ID,
	)
	var i PendingTask
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Date,
		&i.IsDone,
		&i.IsDelete,
		&i.IsFavorite,
	)
	return i, err
}
