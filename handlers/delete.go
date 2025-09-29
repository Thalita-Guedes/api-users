package handlers

import (
	"api-users/repository"
	"api-users/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received on /api/v1/users/{id}")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	err := repository.DeleteUser(r.Context(), id)
	if err != nil {
		utils.ResponseInternalError(w, "Error delete user")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
