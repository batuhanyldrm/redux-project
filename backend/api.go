package main

import (
	"example.com/greetings/models"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	Service *Service
}

func NewApi(service *Service) Api {
	return Api{
		Service: service,
	}
}

func (api *Api) PostTodoHandler(c *fiber.Ctx) error {
	todosDTO := models.TodoDTO{}
	err := c.BodyParser(&todosDTO)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}
	todo, err := api.Service.PostTodos(todosDTO)

	switch err {
	case nil:
		c.JSON(todo)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}
