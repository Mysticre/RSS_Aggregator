package handler

import (
	"net/http"

	"github.com/mysticre/RSS_Aggregator/methods"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	methods.ResponseFor400Error(w, 400, "Something went wrong...")
}
