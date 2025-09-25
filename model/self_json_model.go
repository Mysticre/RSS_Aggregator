package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/mysticre/RSS_Aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

//從sqlc產生的model把它拿過來(參數) 轉成 我們自己生產出來的json model (上面的struct) 要傳出去給clinet使用的

func DatabaseUser(dbUser database.User) User {

	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}

}

func DatabaseFeed(dbFeed database.Feed) Feed {

	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}

}

func DatabaseFeeds(dbFeeds []database.Feed) []Feed {

	feeds := []Feed{} //空箱子切片

	for _, value := range dbFeeds {

		feeds = append(feeds, DatabaseFeed(value)) //記得要再一次存回去空箱子
	}

	return feeds

}

func DatabaseFeedToFollows(dbFeed database.FeedFollow) FeedFollows {

	return FeedFollows{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		UserID:    dbFeed.UserID,
		FeedID:    dbFeed.FeedID,
	}

}

func DatabaseGetFeedToFollows(dbFeedFollow []database.FeedFollow) []FeedFollows {

	feedFollows := []FeedFollows{}

	for _, follows := range dbFeedFollow {

		feedFollows = append(feedFollows, DatabaseFeedToFollows(follows))
	}

	return feedFollows

}
