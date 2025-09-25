-- +goose Up

CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);


/* 
關聯資料庫的用法 子表 可以用 REFERENCE 主表的foreign key 就是user表的id

ON就是 如果主表怎樣 那 子表 就會怎樣的連結 DLETE CASCADE 就是刪除主表的資料 也會順著刪除子表的資料

CASCADE	    刪除/更新 parent 時，子表自動刪除/更新	  eg:  ON DELETE CASCADE
SET NULL	刪除/更新 parent 時，子表外鍵設為 NULL	  eg:  ON DELETE SET NULL
SET DEFAULT	刪除/更新 parent 時，子表外鍵設為預設值	  eg:  ON DELETE SET DEFAULT
RESTRICT	阻止刪除/更新 parent，如果有子表引用	 預設行為 
*/



-- +goose Down

DROP TABLE feeds;