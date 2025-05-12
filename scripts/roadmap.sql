-- Seed data for roadmap
INSERT INTO roadmaps (id, user_id, title, description, created_at)
VALUES 
  (gen_random_uuid(), (SELECT id FROM users LIMIT 1), 'Frontend Foundations', 'Learn core UI and HTML/CSS concepts', NOW()),
  (gen_random_uuid(), (SELECT id FROM users LIMIT 1), 'Backend Bootcamp', 'Learn Go, PostgreSQL and APIs', NOW());
