package main

import (
	"AggregateRSS/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apicfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Println("Error in parsing:CreateFeed")
		respondWithError(w, 400, "Couldn't parse json")
		return
	}

	if params.Name == "" || params.URL == "" {
		respondWithError(w, 400, "Incomplete body")
		return
	}

	createdFeed, err := apicfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		Userid:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, "Couldn't create feed:CreateFeed")
		return
	}

	respondWithJSON(w, 201, mapDbFeedToApiFeed(createdFeed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	fetchedFeeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, "Couldn't fetch feeds:GetFeeds")
		return
	}
	respondWithJSON(w, 201, mapDbFeedsToApiFeeds(fetchedFeeds))
}
