package main

import (
	"encoding/json"
	"net/http"

	"github.com/NishantP43/rssagg/internal/database"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramater struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramater{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{)
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}
