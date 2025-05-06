-- name: CreateUser :one
INSERT INTO users (id, name, email, created_at)
VALUES ($1, $2, $3, $4)

RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: ListAllUsers :many
SELECT * FROM users;
