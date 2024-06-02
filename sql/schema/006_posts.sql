-- +goose Up
CREATE TABLE posts (
  id uuid DEFAULT gen_random_uuid(),
  title TEXT NOT NULL,
  url TEXT NOT NULL,
  description TEXT,
  published_at TIMESTAMP,
  feed_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL default NOW(),
  updated_at TIMESTAMP NOT NULL default NOW(),
  PRIMARY KEY (id),
  FOREIGN KEY (feed_id)
    REFERENCES feeds(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
