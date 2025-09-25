package handler

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

/*從前端拿資料下來*/

type RSSFeed struct {
	Channel struct { //匿名嵌套struct只能存在於RSSFeed
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	}
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

/*

function進去或出來的都是copy
struct在內部或是外部被構造出來






*/

func FetchRSSURLtoFeed(url string) (RSSFeed, error) {

	req := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := req.Get(url)

	if err != nil {
		return RSSFeed{}, err
	}

	defer resp.Body.Close() /*這個body是一個資訊流型態 要記得關閉*/

	dataBytes, err := io.ReadAll(resp.Body) /*資訊流型態 必須用io的reader去讀資料出來 外部讀進來的「原始資料」，通常最終都會先變成[]byte*/

	/*因為 byte 就是最小的資料單位（對應一個 8-bit），所以[]byte可以表示任何資料（二進制、文字、影像、音樂）再根據格式（JSON/XML/Protobuf/自訂協議）用對應的 Unmarshal/Decoder 轉成 struct。*/
	if err != nil {
		return RSSFeed{}, err
	}

	RSSData := RSSFeed{}

	/*雖然RSSData是在函式內部建立的
	但Go的逃逸分析會發現「你把它的位址return 出去了」，
	所以 Go 會把 RSSData 配置到 heap，確保在函式結束後仍然存活。
	外部修改 feed.Title = "X"，會真的改到同一份資料。
	適合 struct 很大（避免複製）或需要後續修改的情境。*/

	err = xml.Unmarshal(dataBytes, &RSSData)
	/*要把拿回還得xml資料 拆掉xml裝在RSSFeed裡面
	這邊傳指標 如果不傳指標，函式就沒辦法修改你外面的變數，只能在自己scope裡做一份拷貝（結果丟了）。
	簡單說就是在一個func內 -> 還要calling一個func完成其他功能再傳出來 基本就是傳進去func2的 就是指標

	*/

	if err != nil {
		return RSSFeed{}, err
	}

	return RSSData, nil

	/*Go在回傳的時候會複製這個struct，
	但因為 RSSData 已經有值，所以外部拿到的是「有值的struct」

	*/

}
