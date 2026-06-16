-- name: CreateUser :one
INSERT INTO users (
    name,
    dob
)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;