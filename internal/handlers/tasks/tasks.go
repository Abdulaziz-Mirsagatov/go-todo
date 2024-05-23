package tasks

import (
	"encoding/json"
	"net/http"

	"github.com/Abdulaziz-Mirsagatov/todo/internal/db"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.Tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, task := range db.Tasks {
		if task.ID.String() == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	json.NewDecoder(r.Body).Decode(&task)
	
	task.ID = uuid.New()
	db.Tasks = append(db.Tasks, task)

	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, task := range db.Tasks {
		if task.ID.String() == id {
			db.Tasks = append(db.Tasks[:i], db.Tasks[i+1:]...)
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedTask db.Task
	json.NewDecoder(r.Body).Decode(&updatedTask)

	for i, task := range db.Tasks {
		if task.ID.String() == id {
			updatedTask.ID = task.ID
			updatedTask.UserID = task.UserID
			db.Tasks[i] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}