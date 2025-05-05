--- name: GetUser :one 
SELECT id, name, email, created_at FROM users WHERE id = $1;

-- name: ListAllUsers :many
SELECT id, name, email, created_at FROM users;

-- name: CreateUser :one
INSERT INTO users (id, name, email, created_at)
VALUES ($1, $2, $3, $4)

RETURNING id, name, email, created_at;