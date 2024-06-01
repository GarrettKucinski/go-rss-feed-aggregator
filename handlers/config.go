package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/garrettkucinski/go-rss-feed-aggregator/internal/database"
)

type Config struct {
	DB *database.Queries
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *Config) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		apiKey := strings.Replace(r.Header.Get("Authorization"), "ApiKey ", "", 1)

		user, err := cfg.DB.GetUserByApiKey(ctx, apiKey)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "could not find user")
			return
		}

		handler(w, r, user)
	}
}
