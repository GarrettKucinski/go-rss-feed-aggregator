-- +goose Up
CREATE TABLE feeds (
  id uuid DEFAULT gen_random_uuid(),
  created_at TIMESTAMP NOT NULL default NOW(),
  updated_at TIMESTAMP NOT NULL default NOW(),
  name TEXT NOT NULL,
  url TEXT NOT NULL,
  user_id UUID NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE,
  CONSTRAINT unique_url UNIQUE(url)
);

-- +goose Down
DROP TABLE feeds;
