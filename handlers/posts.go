package handlers

import (
	"net/http"
	"strconv"

	"github.com/garrettkucinski/go-rss-feed-aggregator/internal/database"
)

func (cfg *Config) HandleGetUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")

	if limitStr == "" {
		limitStr = "15"
	}

	limit, err := strconv.ParseInt(limitStr, 10, 32)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid limit parameter")
		return
	}

	params := database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not get posts for user")
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
