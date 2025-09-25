package handler

import (
	"net/http"

	"github.com/mysticre/RSS_Aggregator/methods"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {

	/*路由器要求的handler格式基本上都是：func(w http.ResponseWriter, r *http.Request)
	  因為：
	  w http.ResponseWriter → 代表「回應管道」，用來回傳資料給前端
	  r *http.Request       → 代表「請求物件」，包含 URL、Header、Body、QueryString 等資訊
	*/

	methods.ResponseWithJson(w, 200, struct{}{}) //struct{}{} 記得是空struct然後直接初始化

}
