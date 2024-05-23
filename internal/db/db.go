package db

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	ID     uuid.UUID `json:"id"`
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

var Users = []User{
	{
		ID:       uuid.New(),
		Username: "user1",
		Password: "password1",
	},
}
var Tasks = []Task{}