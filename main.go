package main

import (
	"api-users/handlers"
	"api-users/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	repository.ConnectDynamoDB()

	repository.CreateTable(repository.Client)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users/", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", handlers.PutUsers).Methods("PUT")
	router.HandleFunc("/api/v1/users/{id}", handlers.DeleteUser).Methods("DELETE")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
