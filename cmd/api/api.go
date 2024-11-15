package api

import (
	"ReactiveX22/code-nest-api/db"
	"ReactiveX22/code-nest-api/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	app := fiber.New()

	app.Get("/api/user", handlers.HandleGetUser)
	app.Get("/api/user/:id", handlers.HandleGetUserByID)
	app.Post("/api/user", handlers.HandleCreateUser)
	app.Post("/api/user/:id", handlers.HandleUpdateUser)
	app.Delete("/api/user/:id", handlers.HandleDeleteUser)

	app.Listen(":3000")
}
