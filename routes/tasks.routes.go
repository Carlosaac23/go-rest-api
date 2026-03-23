package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Carlosaac23/go-rest-api/db"
	"github.com/Carlosaac23/go-rest-api/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	// Create uuid id for task
	id, idErr := uuid.NewV7()
	if idErr != nil {
		w.Write([]byte(idErr.Error()))
	}

	task.ID = id.String()

	createdTask := db.DB.Create(&task)
	createErr := createdTask.Error
	if createErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createErr.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	result := db.DB.Find(&tasks)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, "id = ?", params["id"])

	if task.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, "id = ?", params["id"])

	if task.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	// db.DB.Delete(&task) ---> Change it's state and add deleted_at date
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task successfully deleted"))
}
