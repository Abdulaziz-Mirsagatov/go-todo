package users

import (
	"encoding/json"
	"net/http"

	"github.com/Abdulaziz-Mirsagatov/todo/internal/db"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	usersJson, err := json.Marshal(db.Users)
	if err != nil {
		panic("Failed to convert to JSON")
	}

	w.Write(usersJson)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, user := range db.Users {
		if (user.ID.String() == id) {
			

			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				panic("Failed to convert to JSON")
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("User not found"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic("Failed to decode JSON")
	}

	user.ID = uuid.New()
	db.Users = append(db.Users, user)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		panic("Failed to convert to JSON")
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, user := range db.Users {
		if user.ID.String() == id {
			db.Users = append(db.Users[:i], db.Users[i+1:]...)

			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				panic("Failed to convert to JSON")
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("User not found"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, user := range db.Users {
		if user.ID.String() == id {
			var updatedUser db.User
			err := json.NewDecoder(r.Body).Decode(&updatedUser)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid User"))
				return
			}

			updatedUser.ID = user.ID
			db.Users[i] = updatedUser

			err = json.NewEncoder(w).Encode(updatedUser)
			if err != nil {
				panic("Failed to convert to JSON")
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("User not found"))
}