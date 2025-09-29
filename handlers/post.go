package handlers

import (
	"api-users/models"
	"api-users/repository"
	"api-users/utils"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received on /api/v1/users/")
	w.Header().Set("Content-Type", "application/json")

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Json invalid:", err)
		utils.ResponseError(w, "Request invalid")
		return
	}

	user.ID = utils.GenerateUUID()

	err = repository.CreateUser(r.Context(), user)
	if err != nil {
		log.Println("Repository error:", err)
		utils.ResponseInternalError(w, "Error internal")
		return
	}

	json.NewEncoder(w).Encode(user)
}
