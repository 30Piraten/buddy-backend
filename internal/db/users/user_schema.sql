CREATE TABLE users (
  id UUID PRIMARY KEY,
  name TEXT,
  email TEXT,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE user_roadmaps (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  roadmap_id UUID REFERENCES roadmaps(id),
  started_at TIMESTAMP
);

CREATE TABLE user_checkpoints (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  checkpoint_id UUID REFERENCES checkpoints(id),
  completed_at TIMESTAMP
);

CREATE TABLE events (
  id UUID PRIMARY KEY,
  user_id UUID,
  type TEXT,
  metadata JSONB,
  created_at TIMESTAMP DEFAULT now()
);

