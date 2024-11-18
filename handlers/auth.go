package handlers

import (
	"ReactiveX22/code-nest-api/data"

	"github.com/gofiber/fiber/v2"
)

var sessionTokenCookie = "session_token"

func HandleLogin(c *fiber.Ctx) error {
	loginReq := data.LoginRequest{}

	err := c.BodyParser(&loginReq)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, err, "Invalid Request Body")
	}

	session, err := data.CreateSession(c.Context(), loginReq)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Error Login")
	}

	c.Cookie(&fiber.Cookie{
		Name:     sessionTokenCookie,
		Value:    session.Token,
		Expires:  session.ExpireAt,
		HTTPOnly: true,
		// For development
		// Secure:   true,
	})

	return c.Status(fiber.StatusCreated).JSON(session)
}

func HandleLogout(c *fiber.Ctx) error {
	sessionToken := c.Cookies(sessionTokenCookie)
	if sessionToken == "" {
		return handleError(c, fiber.StatusUnauthorized, nil, "No session token provided")
	}

	err := data.DeleteSession(c.Context(), sessionToken)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Error Logout")
	}

	c.ClearCookie(sessionTokenCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
