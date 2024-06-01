package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/garrettkucinski/go-rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}

func (cfg *Config) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userId, _ := uuid.NewV6()
	res := database.CreateUserParams{
		ID:        userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&res)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not parse user name from request body")
		return
	}

	addedUser, err := cfg.DB.CreateUser(ctx, res)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error adding user")
		return
	}

	respondWithJSON(w, http.StatusOK, addedUser)
}
