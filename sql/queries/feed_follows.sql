-- name: CreateFeedFollows :one

INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)  /*數就是要進入到row的值*/
RETURNING * ;     /*returning 就是回傳所有新增過到row資料 包成struct的方式*/



-- name: GetFollowFeeds :many

SELECT * FROM feed_follows WHERE user_id = $1;

-- name: DeleteFeedFollows :exec    

/*表示只是要執行動作 沒有要返回任何東西*/
DELETE FROM feed_follows WHERE id = $1 AND user_id = $2; /*要是本人才可以刪除 所以加了user_id的條件*/