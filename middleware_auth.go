package main

import (
	"AggregateRSS/internal/auth"
	"AggregateRSS/internal/database"
	"net/http"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 401, "Not Authorised")
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, "No user found with provided api")
			return
		}
		handler(w, r, user)
	}
}
