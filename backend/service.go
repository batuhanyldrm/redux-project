package main

import (
	"strings"
	"time"

	"example.com/greetings/models"
	"github.com/google/uuid"
)

type Service struct {
	repository *Repository
}

func NewService(Repository *Repository) Service {
	return Service{
		repository: Repository,
	}
}

func (service *Service) PostTodos(todosDTO models.TodoDTO) (models.Todo, error) {

	todoCreate := models.Todo{
		ID:          GenerateUUID(8),
		Name:        todosDTO.Name,
		IsCompleted: todosDTO.IsCompleted,
		CreatedAt:   time.Now().UTC().Round(time.Second),
		UpdatedAt:   time.Now().UTC().Round(time.Second),
	}
	CreateTodo, err := service.repository.PostTodos(todoCreate)

	if err != nil {
		return models.Todo{}, err
	}
	return CreateTodo, nil
}

func GenerateUUID(length int) string {
	uuid := uuid.New().String()

	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}
	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}
