-- Seed data for users table
INSERT INTO users (id, name, email, created_at)
VALUES
  ('11111111-1111-1111-1111-111111111111', 'Test User', 'test@buddy.me', CURRENT_TIMESTAMP),
  ('22222222-2222-2222-2222-222222222222', 'Alice Example', 'alice@example.com', CURRENT_TIMESTAMP),
  ('33333333-3333-3333-3333-333333333333', 'Bob Example', 'bob@example.com', CURRENT_TIMESTAMP);
