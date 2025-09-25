-- name: CreateFeed :one

INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6 )  /*數就是要進入到row的值*/
RETURNING * ;     /*returning 就是回傳所有新增過到row資料 包成struct的方式*/


-- name: GetFeeds :many

SELECT * FROM feeds;   /*Many的話就會回傳一個slice裡面裝著所有的Feeds的column內的所有資料*/


-- name: GetNextFeedToFetch :many

SELECT * FROM feeds
ORDER BY last_fetch_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkAsFetchFeed :one

UPDATE feeds 
SET last_fetch_at = now() 
AND updated_at = now()
WHERE id = $1 
RETURNING * ;
