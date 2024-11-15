package handlers

import (
	"ReactiveX22/code-nest-api/data"
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {

	return c.JSON("hello user")
}

func HandleCreateUser(c *fiber.Ctx) error {
	u := data.CreateUserRequest{}
	err := c.BodyParser(&u)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Invalid Request Body")
	}

	createdUser, err := data.CreateUser(c.Context(), u)

	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Error Creating User")
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func HandleGetUserByID(c *fiber.Ctx) error {
	id, err := ParseID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, err, "Invalid user ID format")
	}

	user := &data.User{}

	user, err = data.GetUserByID(c.Context(), id)
	if err != nil {
		if err == data.ErrUserNotFound {
			return handleError(c, fiber.StatusNotFound, err, "User not found")
		}
		return handleError(c, fiber.StatusInternalServerError, err, "Error retrieving user")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func HandleUpdateUser(c *fiber.Ctx) error {
	id, err := ParseID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, err, "Invalid user ID format")
	}

	u := data.UpdateUserRequest{}
	err = c.BodyParser(&u)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Invalid Request Body")
	}

	updatedUser, err := data.UpdateUser(c.Context(), id, u)
	if err != nil {
		if err == data.ErrUserNotFound {
			return handleError(c, fiber.StatusNotFound, err, "User not found")
		}
		return handleError(c, fiber.StatusInternalServerError, err, "Error updating user")
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

func HandleDeleteUser(c *fiber.Ctx) error {
	id, err := ParseID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, err, "Invalid user ID format")
	}

	err = data.DeleteUser(c.Context(), id)
	if err != nil {
		if err == data.ErrUserNotFound {
			return handleError(c, fiber.StatusNotFound, err, "User not found")
		}
		return handleError(c, fiber.StatusInternalServerError, err, "Error deleting user")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
		"userId":  id,
	})
}

// utility functions
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
