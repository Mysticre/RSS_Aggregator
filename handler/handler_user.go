package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mysticre/RSS_Aggregator/internal/auth"
	"github.com/mysticre/RSS_Aggregator/internal/database"
	"github.com/mysticre/RSS_Aggregator/methods"
	"github.com/mysticre/RSS_Aggregator/model"
)

// 這個要放在全域 不能放在main裡面 作用域只在main裡面
type ApiConfig struct {
	DB *database.Queries
}

// dbc 就是「那個正在呼叫方法的apiConfig實體的指標」。也就是說，只有當下這個實體(在記憶體裡)才能使用這個方法
func (dbc *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type paraForJson struct {
		Username string `json:"username"`
	}

	param := paraForJson{}

	decoder := json.NewDecoder(r.Body)
	errD := decoder.Decode(&param)

	if errD != nil {
		methods.ResponseFor400Error(w, 400, "Post body is not valid json")
		return
	}

	user, err := dbc.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      param.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		methods.ResponseFor400Error(w, 500, "Failed to create user")
	}

	methods.ResponseWithJson(w, 200, model.DatabaseUser(user))

}

func HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	methods.ResponseWithJson(w, 200, model.DatabaseUser(user))

}

type authHandler func(w http.ResponseWriter, r *http.Request, u database.User) //刻模板 這才是最終要執行的(要傳入的func)

func (dbc *ApiConfig) MiddlewareAuth(handler authHandler) (h http.HandlerFunc) { //要回傳的是"包裝層" 包進傳入的function

	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			methods.ResponseFor400Error(w, 403, "Invalid API key")
			return
		}

		user, err := dbc.DB.GetUserByAPIkey(r.Context(), apiKey)

		if err != nil {
			methods.ResponseFor400Error(w, 400, fmt.Sprintf("Failed to get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
