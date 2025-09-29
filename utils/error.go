package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseError(w http.ResponseWriter, message string) {
	log.Println("Error decoding request")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func ResponseInternalError(w http.ResponseWriter, message string) {
	log.Println("Error Internal")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func ResponseNotFound(w http.ResponseWriter, message string) {
	log.Println("Error Internal")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
