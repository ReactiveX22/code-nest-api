package handlers

import "github.com/gofiber/fiber/v2"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func HandleGetUser(c *fiber.Ctx) error {

	user := User{
		ID:       1,
		Username: "User01",
		Email:    "user01@email.com",
	}

	return c.JSON(user)
}
