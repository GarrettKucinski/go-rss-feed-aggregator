-- name: CreateRssFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
  VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetAllRssFeeds :many
SELECT * FROM feeds;

-- name: GetAllUserFollows :many
SELECT * from feed_follows where user_id = $1;

-- name: FollowRssFeed :one
INSERT INTO feed_follows(id, created_at, updated_at, feed_id, user_id)
  VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows where feed_id = $1;
