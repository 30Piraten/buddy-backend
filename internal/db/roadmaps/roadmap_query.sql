-- name: CreateRoadmap :one
INSERT INTO roadmaps (user_id, title, description, is_public)
VALUES ($1, $2, $3, $4)
RETURNING *; 

-- name: GetRoadmap :one
SELECT id, user_id, title, description, created_at
FROM roadmaps
WHERE id = $1;

-- name: ListUserRoadmaps :many
SELECT * FROM roadmaps WHERE id = $1 ORDER BY created_at DESC;

-- name: ListAllRoadmaps :many
SELECT * FROM roadmaps;

-- name: UpdateRoadmap :one
UPDATE roadmaps
SET
    title = $2,
    description = $3,
    is_public = $4,
    category = $5,
    tags = $6,
    difficulty = $7
WHERE id = $1
RETURNING *;

-- name: DeleteRoadmap :one
DELETE FROM roadmaps 
WHERE id = $1
RETURNING *; 