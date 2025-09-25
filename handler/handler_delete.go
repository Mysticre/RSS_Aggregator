package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/mysticre/RSS_Aggregator/internal/database"
	"github.com/mysticre/RSS_Aggregator/methods"
)

func (dbc *ApiConfig) HandlerDelete(w http.ResponseWriter, r *http.Request, user database.User) {

	followsIDStr := chi.URLParam(r, "feedFollowID") /*把url的參數取出來(string的狀態)*/

	followsID, err := uuid.Parse(followsIDStr) /*把String轉成UUID*/

	if err != nil {
		methods.ResponseFor400Error(w, 400, err.Error())
		return
	}

	errForDelete := dbc.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     followsID,
		UserID: user.ID,
	})

	if errForDelete != nil {
		methods.ResponseFor400Error(w, 400, fmt.Sprintf("Failed to delete feed follows: %v", errForDelete))
		return
	}

	response := map[string]string{"status": "success"}

	methods.ResponseWithJson(w, 200, response)

}
