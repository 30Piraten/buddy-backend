CREATE TABLE IF NOT EXISTS roadmaps(
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  -- owner_id UUID REFERENCES users(id) ON DELETE CASCADE,
  user_id UUID NOT NULL,
  title TEXT NOT NULL,
  description TEXT,
  is_public BOOLEAN DEFAULT FALSE,
  category TEXT,
  tags TEXT[],
  difficulty TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);