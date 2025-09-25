package handler

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/mysticre/RSS_Aggregator/internal/database"
)

// 啟動型 goroutine 與 定時爬取 功能
func StartScraping(db *database.Queries, concurrency int, timeBetweenRequests time.Duration) {

	log.Printf("Starting scraper with %v gorountines, every %v seconds", concurrency, timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests * time.Second) //每隔(秒)個週期就會往通道送一個時間單位 就是每秒每秒送

	ctx := context.Background() //創造最上層的 root context
	//context 就是用來 控制任務的(goroutine 控制 / API request / metadata 傳遞) 生命週期 與 傳遞控制訊息

	/*
		for迴圈的正確格式

		for [init]; [condition]; [post] { }

		init：迴圈開始時 跑一次在迴圈開始前 只執行一次用來初始化變數、設定狀態
		condition：每圈進來前檢查（省略就等於 true，無窮迴圈）
		body: { } 每圈的內容
		post：每圈 body執行後執行一次

		init →
		[condition?] → body {} → post →
		[condition?] → body {} → post → ...

	*/

	for ; ; <-ticker.C {

		//一次開30個 go rountine搜尋新的feeds
		feeds, err := db.GetNextFeedToFetch(ctx, int32(concurrency))

		if err != nil {
			log.Printf("failed to get next feed to fetch, %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		/*協調 多個 goroutine 完成 的一個同步工具 主線程WaitGroup 就像一個 簽到簿：
		每個 goroutine 進場前要 Add(1) = 簽到
		工作結束要 Done() = 簽退 放在要跑go rountine的function內 +上defer 各自完成後 wg.Done() → 計數器依序遞減到歸0
		主程式 Wait() = 等到簽到簿清空(計數器歸零)，才可以下班  wg.Wait() 放行 → 主程式繼續 */

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(wg, db, feed)
		}

		wg.Wait() //這邊主線程會卡住 等待所有go rountine完成(清空wg記數後)才開始運行
	}
}

func scrapeFeed(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done() //完成該功能後 go rountine 簽退 -1

	feed, err := db.MarkAsFetchFeed(context.Background(), feed.ID)

	if err != nil {
		log.Printf("failed to mark feed as fetched, %v", err)
		return
	}

	RssFeed, err := FetchRSSURLtoFeed(feed.Url)

	if err != nil {
		log.Printf("failed to fetch feed, %v", err)
		return
	}

	for _, item := range RssFeed.Channel.Item { //RssFeed.Channel.Item 是一個 []RSSItem 把每個item的內容拿出來

		log.Printf("found item, %v", item.Title)

	}

}
