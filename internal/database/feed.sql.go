// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feed.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createRssFeed = `-- name: CreateRssFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at
`

type CreateRssFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) CreateRssFeed(ctx context.Context, arg CreateRssFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createRssFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows where feed_id = $1
`

func (q *Queries) DeleteFeedFollow(ctx context.Context, feedID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, feedID)
	return err
}

const followRssFeed = `-- name: FollowRssFeed :one
INSERT INTO feed_follows(id, created_at, updated_at, feed_id, user_id)
  VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at, feed_id, user_id
`

type FollowRssFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) FollowRssFeed(ctx context.Context, arg FollowRssFeedParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, followRssFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedID,
		arg.UserID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedID,
		&i.UserID,
	)
	return i, err
}

const getAllRssFeeds = `-- name: GetAllRssFeeds :many
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at FROM feeds
`

func (q *Queries) GetAllRssFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getAllRssFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.LastFetchedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUserFollows = `-- name: GetAllUserFollows :many
SELECT id, created_at, updated_at, feed_id, user_id from feed_follows where user_id = $1
`

func (q *Queries) GetAllUserFollows(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getAllUserFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNextNFeedsToFetch = `-- name: GetNextNFeedsToFetch :many
SELECT id, url FROM feeds ORDER BY last_fetched_at nulls first LIMIT $1
`

type GetNextNFeedsToFetchRow struct {
	ID  uuid.UUID
	Url string
}

func (q *Queries) GetNextNFeedsToFetch(ctx context.Context, limit int32) ([]GetNextNFeedsToFetchRow, error) {
	rows, err := q.db.QueryContext(ctx, getNextNFeedsToFetch, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNextNFeedsToFetchRow
	for rows.Next() {
		var i GetNextNFeedsToFetchRow
		if err := rows.Scan(&i.ID, &i.Url); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const markFeedFetched = `-- name: MarkFeedFetched :one
UPDATE feeds
  SET last_fetched_at = NOW(), updated_at = NOW()
  WHERE id = $1
  RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at
`

func (q *Queries) MarkFeedFetched(ctx context.Context, id uuid.UUID) (Feed, error) {
	row := q.db.QueryRowContext(ctx, markFeedFetched, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}
