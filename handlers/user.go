package handlers

import "github.com/gofiber/fiber/v2"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func HandleGetUser(c *fiber.Ctx) error {

	return c.JSON("hello user")
}

func HandleCreateUser(c *fiber.Ctx) error {
	u := User{}
	c.BodyParser(&u)

	return c.JSON(u)
}
