package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mysticre/RSS_Aggregator/internal/database"
	"github.com/mysticre/RSS_Aggregator/methods"
	"github.com/mysticre/RSS_Aggregator/model"
)

func (dbc *ApiConfig) HandlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	//從client拿回來的資料 拆掉json 裝在這邊 因為userID已經用middleware拿回來了 所以不用拿前端來的userID
	type ParaForJson struct {
		Feed_id uuid.UUID `json:"feed_id"`
	}

	param := ParaForJson{}

	decoder := json.NewDecoder(r.Body)
	errD := decoder.Decode(&param)

	if errD != nil {
		methods.ResponseFor400Error(w, 400, "Post body is not valid json")
		return
	}

	//存到數據庫後產出一個 儲存結果 壓回 json傳出去到clinet端

	feed_follows, err := dbc.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    param.Feed_id,
	})

	if err != nil {
		methods.ResponseFor400Error(w, 400, fmt.Sprintf("Failed to create feed_follows: %v", err))
		return
	}

	methods.ResponseWithJson(w, 200, model.DatabaseFeedToFollows(feed_follows))

}

func (dbc *ApiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	getFeeds, err := dbc.DB.GetFollowFeeds(r.Context(), user.ID)

	if err != nil {
		methods.ResponseFor400Error(w, 400, fmt.Sprintf("Failed to get feed_follows: %v", err))
		return
	}

	methods.ResponseWithJson(w, 200, model.DatabaseGetFeedToFollows(getFeeds))

}
