package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	repository := NewRepository()
	service := NewService(repository)
	api := NewApi(&service)
	app := SetupApp(&api)
	app.Listen(":3001")

}

func SetupApp(api *Api) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Get("/todos", api.GetTodosHandler)
	app.Post("/todos", api.PostTodosHandler)
	app.Delete("/todos/:id", api.DeleteTodosHandler)
	app.Put("/todos/:id", api.UpdateTodosHandler)
	app.Get("/todos/:id", api.GetTodoHandler)

	return app
}
