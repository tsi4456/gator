-- name: CreatePost :one
INSERT INTO posts
    (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.*, feeds.name as feed_name FROM posts JOIN feed_follows ON posts.feed_id=feed_follows.feed_id JOIN feeds ON feeds.id=feed_follows.feed_id WHERE feed_follows.user_id=$1 ORDER BY published_at DESC LIMIT $2;
