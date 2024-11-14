package api

import (
	"ReactiveX22/code-nest-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	app.Get("/api/user", handlers.HandleGetUser)

	app.Listen(":3000")
}
