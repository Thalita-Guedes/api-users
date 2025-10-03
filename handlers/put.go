package handlers

import (
	"api-users/models"
	"api-users/repository"
	"api-users/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func PutUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received on /api/v1/users/{id}")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Json invalid:", err)
		utils.ResponseError(w, "Request invalid")
		return
	}

	err = repository.UpdateUser(r.Context(), id, user)
	if err != nil {
		log.Println("Error updated  item on  DynamoDB:", err)
		utils.ResponseInternalError(w, "Error internal")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}
