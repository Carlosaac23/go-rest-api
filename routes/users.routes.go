package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Carlosaac23/go-rest-api/db"
	"github.com/Carlosaac23/go-rest-api/helpers"
	"github.com/Carlosaac23/go-rest-api/models"
	"github.com/google/uuid"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	id, idErr := uuid.NewV7()
	if idErr != nil {
		w.Write([]byte(idErr.Error()))
	}

	user.ID = id.String()

	hashedPassword, hashErr := helpers.HashPassword(user.Password)
	if hashErr != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	createdUser := db.DB.Create(&user)
	createErr := createdUser.Error
	if createErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createErr.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
