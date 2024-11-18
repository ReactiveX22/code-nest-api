package handlers

import (
	"ReactiveX22/code-nest-api/data"

	"github.com/gofiber/fiber/v2"
)

func HandleLogin(c *fiber.Ctx) error {
	loginReq := data.LoginRequest{}

	err := c.BodyParser(&loginReq)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Invalid Request Body")
	}

	session, err := data.CreateSession(c.Context(), loginReq)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Error Login")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Expires:  session.ExpireAt,
		HTTPOnly: true,
		// For development
		// Secure:   true,
	})

	return c.Status(fiber.StatusCreated).JSON(session)
}

func HandleLogout(c *fiber.Ctx) error {

	c.ClearCookie("session_token")
	return nil
}
