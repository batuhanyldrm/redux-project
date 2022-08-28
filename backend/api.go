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

func (api *Api) GetTodosHandler(c *fiber.Ctx) error {

	todos, err := api.Service.GetTodos()

	switch err {
	case nil:
		c.JSON(todos)
		c.Status(fiber.StatusOK)

	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) PostTodosHandler(c *fiber.Ctx) error {

	createTodos := models.Todo{}
	err := c.BodyParser(&createTodos)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}
	err = api.Service.PostTodos(createTodos)

	switch err {
	case nil:
		c.Status(fiber.StatusCreated)

	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) DeleteTodosHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	err := api.Service.DeleteTodos(ID)
	switch err {
	case nil:
		c.Status(fiber.StatusNoContent)

	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) UpdateTodosHandler(c *fiber.Ctx) error {

	ID := c.Params("id")
	todo := models.TodoDTO{}
	err := c.BodyParser(&todo)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	updatedTodo, err := api.Service.UpdateTodos(todo, ID)

	switch err {
	case nil:
		//response
		c.JSON(updatedTodo)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) GetTodoHandler(c *fiber.Ctx) error {

	ID := c.Params("id")
	todo, err := api.Service.GetTodo(ID)

	switch err {
	case nil:
		//response
		c.JSON(todo)
		c.Status(fiber.StatusOK)
	/* case TodoNotFound:
	c.Status(fiber.StatusNotFound) */
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}
