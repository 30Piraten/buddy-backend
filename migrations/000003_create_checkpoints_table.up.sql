CREATE TABLE IF NOT EXISTS checkpoints(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  roadmap_id UUID NOT NULL,
  title TEXT NOT NULL,
  description TEXT,
  position INT NOT NULL,
  type TEXT CHECK (type IN ('LEARNING', 'PRACTICE', 'ASSESSMENT')) NOT NULL,
  status TEXT CHECK (status IN ('PENDING', 'IN_PROGRESS', 'COMPLETED')) NOT NULL, 
  estimated_time INT, 
  reward_points INT,
  created_at TIMESTAMPTZ DEFAULT NOW()
)