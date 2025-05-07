-- name: CreateCheckpoint :one
INSERT INTO checkpoints (roadmap_id, title, description, position, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *; 

-- name: GetCheckpoint :one
SELECT * FROM checkpoints WHERE id = $1; 

-- name: ListCheckpoints :many
SELECT * FROM checkpoints WHERE roadmap_id = $1 ORDER BY position ASC; 