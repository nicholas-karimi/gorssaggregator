package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nicholas-karimi/gorssaggregator/internal/auth"
	"github.com/nicholas-karimi/gorssaggregator/internal/database"
)

func (apiCfg *APIConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing JSON:", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error creating user:", err))
		return
	}

	// respondWithJSON(w, 200, user)
	respondWithJSON(w, 201, databaseUserToUser(user))

}


func (apiCfg *APIConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth with error %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error getting user:", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}