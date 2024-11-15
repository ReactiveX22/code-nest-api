package handlers

import (
	"ReactiveX22/code-nest-api/data"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {

	return c.JSON("hello user")
}

func HandleCreateUser(c *fiber.Ctx) error {
	u := data.CreateUserRequest{}
	err := c.BodyParser(&u)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	createdUser, err := data.CreateUser(c.Context(), u)

	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error Creating User",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}
