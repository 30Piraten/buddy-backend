CREATE TABLE checkpoints (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  roadmap_id UUID NOT NULL,
  title TEXT NOT NULL,
  description TEXT,
  position INT NOT NULL,
  type TEXT CHECK (type IN ('LEARNING', 'PRACTICE', 'ASSESSMENT')) NOT NULL,
  status TEXT CHECK (status IN ('PENDING', 'IN_PROGRESS', 'COMPLETED')) NOT NULL, 
  estimated_time INT, 
  reward_points INT,
  created_at TIMESTAMPZ DEFAULT NOW()
);

CREATE INDEX idx_checkpoint_roadmap_position ON checkpoints (roadmap_id, position);

CREATE TABLE user_checkpoints (
  user_id UUID NOT NULL,
  checkpoint_id UUID NOT NULL REFERENCES checkpoints(id) ON DELETE CASCADE,
  status TEXT CHECK (status IN ('PENDING', 'IN_PROGRESS', 'COMPLETED')) NOT NULL, 
  PRIMARY key (user_id, checkpoint_id)
); 