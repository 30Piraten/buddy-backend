CREATE TABLE checkpoints (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  roadmap_id UUID REFERENCES roadmaps(id) ON DELETE CASCADE,
  title TEXT NOT NULL,
  description TEXT,
  position INT NOT NULL,
  type TEXT NOT NULL,
  status TEXT DEFAULT 'pending', 
  estimated_time INT, 
  reward_points INT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);