package handler

import (
	"net/http"

	"github.com/mysticre/RSS_Aggregator/methods"
	"github.com/mysticre/RSS_Aggregator/model"
)

func (dbc *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := dbc.DB.GetFeeds(r.Context())
	if err != nil {
		methods.ResponseFor400Error(w, 500, "Failed to get feeds")
		return
	}
	methods.ResponseWithJson(w, 200, model.DatabaseFeeds(feeds))
}
