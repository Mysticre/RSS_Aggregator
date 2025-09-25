-- name: GetPosts :one

INSERT INTO posts (
    id,
    created_at,
    updated_at,
    title,
    description,
    url,
    feed_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)

RETURNING * ;


-- name: JoinAndGetPostsAsUser :one

SELECT post* FROM posts
Join feed_follows ON Posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.updated_at DESC 
LIMIT $2 ; 

