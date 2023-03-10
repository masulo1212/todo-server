// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: bin.sql

package db

import (
	"context"
)

const createBin = `-- name: CreateBin :one
INSERT INTO bin_tasks (
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

type CreateBinParams struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	IsDone      bool   `json:"is_done"`
	IsDelete    bool   `json:"is_delete"`
	IsFavorite  bool   `json:"is_favorite"`
}

func (q *Queries) CreateBin(ctx context.Context, arg CreateBinParams) (BinTask, error) {
	row := q.db.QueryRowContext(ctx, createBin,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Date,
		arg.IsDone,
		arg.IsDelete,
		arg.IsFavorite,
	)
	var i BinTask
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

const deleteBin = `-- name: DeleteBin :exec
DELETE FROM bin_tasks WHERE id = $1
`

func (q *Queries) DeleteBin(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteBin, id)
	return err
}

const getAllBin = `-- name: GetAllBin :many
SELECT id, title, description, date, is_done, is_delete, is_favorite FROM bin_tasks
ORDER BY date
`

func (q *Queries) GetAllBin(ctx context.Context) ([]BinTask, error) {
	rows, err := q.db.QueryContext(ctx, getAllBin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BinTask{}
	for rows.Next() {
		var i BinTask
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

const getBin = `-- name: GetBin :one
SELECT id, title, description, date, is_done, is_delete, is_favorite FROM bin_tasks
WHERE id = $1
`

func (q *Queries) GetBin(ctx context.Context, id string) (BinTask, error) {
	row := q.db.QueryRowContext(ctx, getBin, id)
	var i BinTask
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
