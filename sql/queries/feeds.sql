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
--

-- name: GetFeeds :many
SELECT * FROM feeds;
--

-- name: GetFeedByUrl :one
SELECT * 
FROM feeds
WHERE url=$1;
--

-- name: GetFeedFollowsByUrl :many
SELECT f.name, u.name, ff.id
FROM feed_follows ff
INNER JOIN feeds f
ON f.id = ff.feed_id
INNER JOIN users u
ON u.id = ff.user_id
WHERE u.id=$1;
--

-- name: DeleteFeedFollowsForUser :one
DELETE FROM feed_follows ff
WHERE ff.feed_id IN (SELECT id FROM feeds WHERE url = $1)
AND ff.user_id=$2
RETURNING *;
--

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
--

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1;
--

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at=$1,
    updated_at=$2
WHERE id=$3
RETURNING *;
--

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
--