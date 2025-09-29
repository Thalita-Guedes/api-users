package handlers

import (
	"api-users/repository"
	"api-users/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received on /api/v1/users/{id}")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	user, err := repository.GetUser(r.Context(), id)
	if err != nil {
		log.Println("Error fetching item on DynamoDB:", err)
		utils.ResponseInternalError(w, "Internal error")
		return
	}

	if user == nil {
		utils.ResponseNotFound(w, "User not found")
		return
	}

	json.NewEncoder(w).Encode(user)
}
