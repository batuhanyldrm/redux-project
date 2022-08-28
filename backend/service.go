package main

import (
	"strings"

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

func (service *Service) PostTodos(todo models.Todo) error {

	todo.ID = GenerateUUID(8)

	err := service.Repository.PostTodos(todo)

	if err != nil {
		return err
	}

	return nil
}

func (service *Service) DeleteTodos(todoId string) error {

	err := service.Repository.DeleteTodos(todoId)

	if err != nil {
		return err
	}

	return nil
}

func (service *Service) UpdateTodos(todo models.TodoDTO, ID string) (models.Todo, error) {

	updatedTodo, err := service.Repository.UpdateTodos(todo, ID)

	if err != nil {
		return models.Todo{}, err
	}

	return updatedTodo, nil
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
