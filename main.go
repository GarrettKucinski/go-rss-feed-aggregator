package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/garrettkucinski/go-rss-feed-aggregator/handlers"
	"github.com/garrettkucinski/go-rss-feed-aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))

	if err != nil {
		log.Fatal("unable to connect to db")
	}

	dbQueries := database.New(db)

	cfg := handlers.Config{
		DB: dbQueries,
	}

	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handlers.HandleHealthCheck)
	mux.HandleFunc("GET /v1/err", handlers.HandleErrors)

	mux.HandleFunc("GET /v1/users", cfg.MiddlewareAuth(handlers.HandleGetUser))
	mux.HandleFunc("GET /v1/feeds", cfg.HandleGetAllFeeds)
	mux.HandleFunc("GET /v1/feed_follows", cfg.MiddlewareAuth(cfg.HandleGetFollowsForUser))

	mux.HandleFunc("POST /v1/users", cfg.HandleCreateUser)
	mux.HandleFunc("POST /v1/feeds", cfg.MiddlewareAuth(cfg.HandleCreateFeed))
	mux.HandleFunc("POST /v1/feed_follows", cfg.MiddlewareAuth(cfg.HandleCreateFeedFollow))

	mux.HandleFunc("DELETE /v1/feed_follows", cfg.MiddlewareAuth(cfg.HandleDeleteFeedFollow))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go cfg.RssFeedWorker(10)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
