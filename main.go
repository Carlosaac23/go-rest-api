package main

import (
	"log"
	"net/http"

	"github.com/Carlosaac23/go-rest-api/db"
	"github.com/Carlosaac23/go-rest-api/models"
	"github.com/Carlosaac23/go-rest-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	log.Print("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
