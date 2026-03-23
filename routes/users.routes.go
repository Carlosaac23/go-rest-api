package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Carlosaac23/go-rest-api/db"
	"github.com/Carlosaac23/go-rest-api/helpers"
	"github.com/Carlosaac23/go-rest-api/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Create uuid id for user
	id, idErr := uuid.NewV7()
	if idErr != nil {
		w.Write([]byte(idErr.Error()))
	}

	user.ID = id.String()

	// Hash user password
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
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	result := db.DB.Find(&users)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, "id = ?", params["id"])

	if user.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, "id = ?", params["id"])

	if user.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	// db.DB.Delete(&user) ---> Change it's state and add deleted_at date
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User successfully deleted"))
}
