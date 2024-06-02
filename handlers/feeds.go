package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/garrettkucinski/go-rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *Config) HandleGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	feeds, err := cfg.DB.GetAllRssFeeds(ctx)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}

func (cfg *Config) HandleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	ctx := context.Background()

	feedId, _ := uuid.NewV6()
	feed := database.CreateRssFeedParams{
		ID:        feedId,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&feed)

	feedRes, err := cfg.DB.CreateRssFeed(ctx, feed)

	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "could not create feed")
		return
	}

	follow := database.FollowRssFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	followedFeed, err := cfg.DB.FollowRssFeed(r.Context(), follow)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error following created feed")
		return
	}

	respondWithJSON(w, http.StatusOK, struct {
		Feed       database.Feed
		FeedFollow database.FeedFollow
	}{
		Feed:       feedRes,
		FeedFollow: followedFeed,
	})
}

func (cfg *Config) HandleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feed := database.FollowRssFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&feed)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not decode feed id")
		return
	}

	followedFeed, err := cfg.DB.FollowRssFeed(r.Context(), feed)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error following feed")
		return
	}

	respondWithJSON(w, http.StatusOK, followedFeed)
}

func (cfg *Config) HandleGetFollowsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	follows, err := cfg.DB.GetAllUserFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not get follows for user")
		return
	}

	respondWithJSON(w, http.StatusOK, follows)
}

func (cfg *Config) HandleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		FeedId uuid.UUID
	}

	feed := params{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&feed)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not parse feed Id")
		return
	}

	deleteErr := cfg.DB.DeleteFeedFollow(r.Context(), feed.FeedId)

	if deleteErr != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to delete follow")
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
