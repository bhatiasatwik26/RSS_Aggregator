package main

import (
	"AggregateRSS/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		UserName string `json:name`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	fmt.Println(params)
	if err != nil {
		respondWithError(w, 400, "Invalid user input")
		return
	}
	createdUser, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		Name:      params.UserName,
	})

	if err != nil {
		fmt.Println("error: ", err)
		respondWithError(w, 400, "Coudln't create the user")
		return
	}
	respondWithJSON(w, 201, mapDbUserToApiUser(createdUser))
}

func (apiConfig *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, 200, mapDbUserToApiUser(user))
}
