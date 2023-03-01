// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: fav.sql

package db

import (
	"context"
	"database/sql"
)

const createFav = `-- name: CreateFav :one
INSERT INTO fav_tasks (
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

type CreateFavParams struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	IsDone      bool   `json:"is_done"`
	IsDelete    bool   `json:"is_delete"`
	IsFavorite  bool   `json:"is_favorite"`
}

func (q *Queries) CreateFav(ctx context.Context, arg CreateFavParams) (FavTask, error) {
	row := q.db.QueryRowContext(ctx, createFav,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Date,
		arg.IsDone,
		arg.IsDelete,
		arg.IsFavorite,
	)
	var i FavTask
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

const deleteFav = `-- name: DeleteFav :exec
DELETE FROM fav_tasks WHERE id = $1
`

func (q *Queries) DeleteFav(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteFav, id)
	return err
}

const getAllFav = `-- name: GetAllFav :many
SELECT id, title, description, date, is_done, is_delete, is_favorite FROM fav_tasks
ORDER BY date
`

func (q *Queries) GetAllFav(ctx context.Context) ([]FavTask, error) {
	rows, err := q.db.QueryContext(ctx, getAllFav)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FavTask{}
	for rows.Next() {
		var i FavTask
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

const getFav = `-- name: GetFav :one
SELECT id, title, description, date, is_done, is_delete, is_favorite FROM fav_tasks
WHERE id = $1
`

func (q *Queries) GetFav(ctx context.Context, id string) (FavTask, error) {
	row := q.db.QueryRowContext(ctx, getFav, id)
	var i FavTask
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

const updateFav = `-- name: UpdateFav :one
UPDATE fav_tasks
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

type UpdateFavParams struct {
	IsDone      bool           `json:"is_done"`
	IsDelete    bool           `json:"is_delete"`
	IsFavorite  bool           `json:"is_favorite"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	Date        sql.NullString `json:"date"`
	ID          string         `json:"id"`
}

func (q *Queries) UpdateFav(ctx context.Context, arg UpdateFavParams) (FavTask, error) {
	row := q.db.QueryRowContext(ctx, updateFav,
		arg.IsDone,
		arg.IsDelete,
		arg.IsFavorite,
		arg.Title,
		arg.Description,
		arg.Date,
		arg.ID,
	)
	var i FavTask
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