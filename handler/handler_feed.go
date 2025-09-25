package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mysticre/RSS_Aggregator/internal/database"
	"github.com/mysticre/RSS_Aggregator/methods"
	"github.com/mysticre/RSS_Aggregator/model"
)

func (dbc *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type ParaForJson struct {
		Username string `json:"username"`
		URL      string `json:"url"`
	}

	param := ParaForJson{}

	decoder := json.NewDecoder(r.Body)
	errD := decoder.Decode(&param)

	if errD != nil {
		methods.ResponseFor400Error(w, 400, "Post body is not valid json")
		return
	}

	feed, err := dbc.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      param.Username,
		Url:       param.URL,
		UserID:    user.ID,
	})

	if err != nil {
		methods.ResponseFor400Error(w, 500, "Failed to create Feed")
	}

	methods.ResponseWithJson(w, 200, model.DatabaseFeed(feed))

}
