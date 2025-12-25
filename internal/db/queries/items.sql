-- name: CreateItem :one
INSERT INTO todo_items (
    title,
    description,
    image,
    status
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateItem :one
UPDATE todo_items
SET
    title = COALESCE(sqlc.narg(title), title),
    description = COALESCE(sqlc.narg(description), description),
    image = COALESCE(sqlc.narg(image), image),
    status = COALESCE(sqlc.narg(status), status)
WHERE
    id = sqlc.arg(id)
    AND deleted_at IS NULL
RETURNING *;

-- name: GetItemDetail :one
SELECT *
FROM todo_items
WHERE id = $1;

-- name: DeleteItem :one
DELETE FROM todo_items
WHERE
    id = $1
    AND deleted_at IS NULL
RETURNING *;

-- name: RestoreItem :one
UPDATE todo_items
SET status = 1
WHERE
    id = sqlc.arg(id)
    AND status = 3
RETURNING *;