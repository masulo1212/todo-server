-- name: CreatePending :one
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
RETURNING *;

-- name: GetPending :one
SELECT * FROM pending_tasks
WHERE id = $1;


-- name: GetAllPending :many
SELECT * FROM pending_tasks
ORDER BY date;


-- name: UpdatePending :one
UPDATE pending_tasks
SET
   title = COALESCE(sqlc.narg(title),title),
  description = COALESCE(sqlc.narg(description),description),
  date = COALESCE(sqlc.narg(date),date),
  is_done = $1,
  is_delete = $2,
  is_favorite = $3
WHERE
   id = sqlc.arg(id)
RETURNING *;


-- name: DeletePending :exec
DELETE FROM pending_tasks WHERE id = $1;