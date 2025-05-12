-- Seed data for users table
INSERT INTO users (id, name, email, created_at)
VALUES
  (gen_random_uuid(), 'Test User', 'test@buddy.me', CURRENT_TIMESTAMP),
  (gen_random_uuid(), 'Alice Example', 'alice@example.com', CURRENT_TIMESTAMP),
  (gen_random_uuid(), 'Bob Example', 'bob@example.com', CURRENT_TIMESTAMP);

