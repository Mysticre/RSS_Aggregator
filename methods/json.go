package methods //寫package main 就可以視為同一群組 不用再另外import就可以共用一些方法

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseFor400Error(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Printf("Respoding with 5xx error: %s", msg)
	}

	type errResponse struct {
		Err string `json:"error"`
		//json的key tag  然後 1.參數到導出要大寫 2.中間不要有空格 struct轉json的key
	}

	ResponseWithJson(w, code, errResponse{Err: msg}) //專門處理json回傳 只是把interace換成error msg

}

// 專門處理 + 壓成json資料 + 用w把資料寫回去
func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload) //把Go的資料結構（struct、map、slice 等）轉成JSON格式 Go偏向「先產生 byte，再決定用途」

	if err != nil {
		log.Printf("Failed to marshal JSON: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

	/*ResponseWriter = 寫回應給前端的接口
	  你可以：
	  設定 header (Header())
	  設定狀態碼 (WriteHeader)
	  寫入回應內容 (Write)
	  搭配 encoding/json 可以直接回傳 JSON

	  map用type之後就可以當作"方法"的接收者 只有type定義的可以當作方法的接收者

	*/

}
