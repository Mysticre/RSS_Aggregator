package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(h http.Header) (string, error) {

	/*func 傳入型別 是否要加指標*
	通常取決於兩個因素：型別本身是 值-型別 還是 引用-型別 + 你要不要修改原本的資料
	基本型別（int, bool, float), struct（如 http.Request）- 值型別
	map（如 http.Header), slice, interface - 引用型別
	*/

	val := h.Get("Authorization")

	if val == " " {
		return " ", errors.New("API key is not found in header")
	}

	/*會取出 "Basic <Base64字串>" 或 "Bearer <token字串>"
	"Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ=="              Basic認證
	"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."  JWT Bearer認證
	*/

	parts := strings.Split(val, " ")

	if len(parts) != 2 {
		return " ", errors.New("Malformed API key in header")
	}

	return parts[1], nil

}
