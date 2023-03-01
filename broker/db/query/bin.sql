-- name: CreateBin :one
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
RETURNING *;

-- name: GetBin :one
SELECT * FROM bin_tasks
WHERE id = $1;


-- name: GetAllBin :many
SELECT * FROM bin_tasks
ORDER BY date;


-- name: DeleteBin :exec
DELETE FROM bin_tasks WHERE id = $1;

-- nemeL DeleteAllBin :exec
DELETE FROM bin_tasks;