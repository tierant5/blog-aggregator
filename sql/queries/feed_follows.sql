
-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT 
    ff.id AS feed_follow_id,
    users.name AS user_name,
    feeds.name AS feed_name
FROM feed_follows ff
INNER JOIN users ON ff.user_id = users.id
INNER JOIN feeds ON ff.feed_id = feeds.id
WHERE ff.user_id = $1;
