-- Seed checkpoints linked to each roadmap
INSERT INTO checkpoints (
  id, roadmap_id, title, description, position, type, status, estimated_time, reward_points, created_at
)
VALUES
  (gen_random_uuid(), gen_random_uuid(), 'Basic HTML', 'Build a semantic HTML page', 1, 'LEARNING', 'PENDING', 30, 10, NOW()),
  (gen_random_uuid(), gen_random_uuid(), 'CSS Flexbox', 'Layout with Flexbox', 2, 'PRACTICE', 'PENDING', 45, 15, NOW()),
  (gen_random_uuid(), gen_random_uuid(), 'Hello, Go', 'Write your first Go program', 1, 'LEARNING', 'PENDING', 25, 10, NOW()),
  (gen_random_uuid(), gen_random_uuid(), 'REST API', 'Create a RESTful endpoint', 2, 'ASSESSMENT', 'IN_PROGRESS', 60, 25, NOW());


-- Seed data for user_checkpoints
-- INSERT INTO user_checkpoints (id, roadmap_id, status)
-- VALUES
--     (gen_random_uuid(), (SELECT id FROM roadmaps LIMIT 1), "COMPLETED");
--     (gen_random_uuid(), (SELECT id FROM roadmaps LIMIT 1), "IN_PROGRESS");