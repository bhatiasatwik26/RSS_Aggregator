package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dataInBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshall response", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(dataInBytes)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Error 5xx:", msg)
	}
	respondWithJSON(w, code, struct {
		Err string `json:"error"`
	}{Err: msg})
}
