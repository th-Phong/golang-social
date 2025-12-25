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
    image = COALESCE(sqlc.narg(image), image)
WHERE
    id = sqlc.arg(id)
RETURNING *;


