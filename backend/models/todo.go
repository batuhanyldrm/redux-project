package models

import "time"

type Todo struct {
	ID          string    `json:"id" bson:"id"`
	Name        string    `json:"name" bson:"name"`
	IsCompleted bool      `json:"isCompleted" bson:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}

type TodoDTO struct {
	Name        string    `json:"name" bson:"name"`
	IsCompleted bool      `json:"isCompleted" bson:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}
