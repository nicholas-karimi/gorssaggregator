package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	/* Responds with an error message */

	if code > 499 {
		log.Println("Responding with 5XX error: ", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: message,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	/* Responds with arbitrary JSON */
	dat, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to Marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
