-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByUrl :one
SELECT * 
FROM feeds
WHERE url=$1;

-- name: GetFeedFollowsForUser :many
SELECT f.name, u.name, ff.id
FROM feed_follows ff
INNER JOIN feeds f
ON f.id = ff.feed_id
INNER JOIN users u
ON u.id = ff.user_id
WHERE u.id=$1;



-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)

SELECT 
    inserted_feed_follow.*,
    f.name AS feed_name,
    u.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds f
ON f.id = inserted_feed_follow.feed_id
INNER JOIN users u
ON u.id = inserted_feed_follow.user_id;