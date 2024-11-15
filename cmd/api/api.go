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

	// users
	app.Get("/api/users", handlers.HandleGetUser)
	app.Get("/api/users/:id", handlers.HandleGetUserByID)
	app.Post("/api/users", handlers.HandleCreateUser)
	app.Patch("/api/users/:id", handlers.HandleUpdateUser)
	app.Delete("/api/users/:id", handlers.HandleDeleteUser)

	// posts
	app.Get("/api/posts", handlers.HandleGetPost)
	app.Get("/api/posts/:id", handlers.HandleGetPostByID)
	app.Post("/api/posts", handlers.HandleCreatePost)
	app.Patch("/api/posts/:id", handlers.HandleUpdatePost)
	app.Delete("/api/posts/:id", handlers.HandleDeletePost)

	app.Listen(":3000")
}
