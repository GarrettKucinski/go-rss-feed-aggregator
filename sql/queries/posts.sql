-- name: CreateRssPost :one
INSERT INTO posts (id, feed_id, title, url, description, published_at, created_at, updated_at)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetPostsByUser :many
SELECT p.* FROM posts AS p
RIGHT JOIN feed_follows AS ff
ON p.feed_id = ff.feed_id
WHERE ff.user_id = $1
LIMIT $2;
