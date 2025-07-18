-- name: CreateFeed :one
INSERT INTO feeds
    (id, created_at, updated_at, name, url, user_id)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name AS username
FROM feeds JOIN users ON feeds.user_id=users.id;

-- name: GetFeedForURL :one
SELECT * FROM feeds WHERE url=$1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at=now(), updated_at=now() WHERE id=$1;

-- name: GetNextFeedToFetch :one
SELECT id, url FROM feeds ORDER BY last_fetched_at NULLS FIRST LIMIT 1;
