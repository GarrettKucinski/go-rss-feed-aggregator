-- +goose Up
CREATE TABLE feed_follows (
  id uuid DEFAULT gen_random_uuid(),
  created_at TIMESTAMP NOT NULL default NOW(),
  updated_at TIMESTAMP NOT NULL default NOW(),
  feed_id UUID NOT NULL,
  user_id UUID NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (feed_id) REFERENCES feeds(id)
    ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE,
  CONSTRAINT unique_feed_user UNIQUE(feed_id, user_id)
);

-- +goose Down
DROP TABLE feed_follows;
