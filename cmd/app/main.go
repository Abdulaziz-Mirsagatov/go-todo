package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abdulaziz-Mirsagatov/todo/internal/handlers/hello"
	"github.com/Abdulaziz-Mirsagatov/todo/internal/handlers/tasks"
	"github.com/Abdulaziz-Mirsagatov/todo/internal/handlers/users"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	usersRouter := mux.PathPrefix("/users").Subrouter()
	tasksRouter := mux.PathPrefix("/tasks").Subrouter()

	mux.HandleFunc("/hello", hello.HelloHandler).Methods("GET")

	usersRouter.HandleFunc("", users.GetUsers).Methods("GET")
	usersRouter.HandleFunc("/{id}", users.GetUser).Methods("GET")
	usersRouter.HandleFunc("/create", users.CreateUser).Methods("POST")
	usersRouter.HandleFunc("/{id}", users.DeleteUser).Methods("DELETE")
	usersRouter.HandleFunc("/{id}", users.UpdateUser).Methods("PUT")

	tasksRouter.HandleFunc("", tasks.GetTasks).Methods("GET")
	tasksRouter.HandleFunc("/{id}", tasks.GetTask).Methods("GET")
	tasksRouter.HandleFunc("/create", tasks.CreateTask).Methods("POST")
	tasksRouter.HandleFunc("/{id}", tasks.DeleteTask).Methods("DELETE")
	tasksRouter.HandleFunc("/{id}", tasks.UpdateTask).Methods("PUT")

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Server connection failed:", err)
	}
}