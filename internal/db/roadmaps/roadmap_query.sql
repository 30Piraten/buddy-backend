-- name: CreateRoadmap :one
INSERT INTO roadmaps (owner_id, title, description, is_public)
VALUES ($1, $2, $3, $4)
RETURNING *; 

-- name: GetRoadmap :one 
SELECT * FROM roadmaps WHERE id = $1;

-- name: ListUserRoadmaps :many
SELECT * FROM roadmaps WHERE owner_id = $1 ORDER BY created_at DESC;