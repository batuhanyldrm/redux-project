package main

import (
	"strings"
	"time"

	"example.com/greetings/models"
	"github.com/google/uuid"
)

type Service struct {
	Repository *Repository
}

func NewService(Repository *Repository) Service {
	return Service{
		Repository: Repository,
	}
}

func (service *Service) GetTodos() ([]models.Todo, error) {

	todos, err := service.Repository.GetTodos()

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (service *Service) PostTodos(todoDTO models.TodoDTO) *models.Todo {

	todo := models.Todo{}
	todo.ID = GenerateUUID(8)
	todo.CreatedAt = time.Now().UTC().Round(time.Second)
	todo.UpdatedAt = time.Now().UTC().Round(time.Second)
	todo.IsCompleted = false
	todo.Name = todoDTO.Name

	err := service.Repository.PostTodos(todo)
	if err != nil {
		return nil
	}

	return &todo
}

func (service *Service) DeleteTodos(todoId string) error {

	err := service.Repository.DeleteTodos(todoId)

	if err != nil {
		return err
	}

	return nil
}

func (service *Service) UpdateTodos(todoDTO models.TodoDTO, ID string) (*models.Todo, error) {

	todo, err := service.Repository.GetTodo(ID)

	if err != nil {
		return nil, err
	}

	todo.Name = todoDTO.Name
	todo.UpdatedAt = time.Now().UTC().Round(time.Second)

	err = service.Repository.UpdateTodos(todo, ID)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (service *Service) GetTodo(ID string) (models.Todo, error) {

	updatedTodo, err := service.Repository.GetTodo(ID)

	if err != nil {
		return models.Todo{}, err
	}

	return updatedTodo, nil
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
