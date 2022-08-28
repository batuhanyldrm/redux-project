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
	app.Listen(":3002")
}

func SetupApp(api *Api) *fiber.App {
	app := fiber.New()

	app.Post("/todos", api.PostTodoHandler)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	return app
}
