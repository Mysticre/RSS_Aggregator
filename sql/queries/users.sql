-- name: CreateUser :one

INSERT INTO users (id, name, created_at, updated_at, api_key)
VALUES ($1, $2, $3, $4, encode(digest(random()::text, 'sha256'), 'hex'))  /*數就是要進入到row的值*/
RETURNING * ;     /*returning 就是回傳所有新增過到row資料 包成struct的方式*/



-- name: GetUserByAPIkey :one
SELECT * FROM users WHERE api_key = $1;