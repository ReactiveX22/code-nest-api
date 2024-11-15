package handlers

import (
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, status int, err error, message string) error {
	log.Printf("%s: %v", message, err)
	return c.Status(status).JSON(fiber.Map{
		"error":   message,
		"message": err.Error(),
	})
}

func ParseID(c *fiber.Ctx) (int, error) {
	params := c.Params("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return 0, errors.New("invalid user ID format")
	}
	return id, nil
}
