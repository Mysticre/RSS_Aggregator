package main

import (
	_ "context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	/*每一個「任務」（例如一次HTTP請求、一個DB查詢、一個背景工作都會有一個「上下文」這個上下文會帶著「控制訊號 + 請求資料」，傳遞給整個任務鏈路👉 所以它就叫 context（上下文、環境）。*/

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //是 Go語言的一個PostgreSQL驅動程式 (driver) _ 前面+底線表示 只引入driver不直接使用程式碼
	"github.com/mysticre/RSS_Aggregator/handler"
	"github.com/mysticre/RSS_Aggregator/internal/database"
	"github.com/rs/cors"
)

func main() {

	//連線資料庫需要用到的struct

	godotenv.Load(".env")

	port := os.Getenv("PORT") //從key去取得資訊 (就是 PORT變數)

	if port == "" {
		log.Fatal("Port is ont found in env")
	}

	fmt.Printf("Port %s is open.", port)

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in env")
	}

	//開啟資料庫
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Unable to connect to database: %v", err)
	}

	//DB連線後要產生要執行寫入的context
	dbConn := database.New(conn)

	dbConfig := handler.ApiConfig{
		DB: dbConn,
	}

	go handler.StartScraping(dbConn, 10, time.Minute) //啟動另外一個線程去撈資料 持續背景型 就不用控制了
	/*
		持續背景型 → go func() 就好
		一次性並行任務 需要「結果」回傳給主線程（要用 channel） → 搭配 WaitGroup 或 channel
		需要控制生命周期 → 搭配 context

	*/

	router := chi.NewRouter() //主router

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler(router)

	log.Printf("Server is running on port %s", port)

	v1Router := chi.NewRouter() //副v1 router
	v1Router.Get("/health", handler.HandlerReadiness)
	v1Router.Get("/err", handler.HandlerError)

	v1Router.Post("/users", dbConfig.HandlerCreateUser)
	v1Router.Get("/users", dbConfig.MiddlewareAuth(handler.HandlerGetUser))

	v1Router.Post("/feed", dbConfig.MiddlewareAuth(dbConfig.HandlerCreateFeed))
	v1Router.Get("/feeds", dbConfig.HandlerGetFeeds)

	v1Router.Post("/feed_follows", dbConfig.MiddlewareAuth(dbConfig.HandlerCreateFeedFollows))
	v1Router.Get("/feed_follows", dbConfig.MiddlewareAuth(dbConfig.HandlerGetFeedFollows))

	v1Router.Delete("/feed_follows/{feedFollowID}", dbConfig.MiddlewareAuth(dbConfig.HandlerDelete))

	router.Mount("/v1", v1Router) //把v1副router掛載到主router上面    就是:8080/v1/health

	er := srv.ListenAndServe()
	if er != nil {
		log.Fatal(er)
	}

}
