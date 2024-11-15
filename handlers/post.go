package handlers

import (
	"ReactiveX22/code-nest-api/data"

	"github.com/gofiber/fiber/v2"
)

func HandleGetPost(c *fiber.Ctx) error {
	posts, err := data.GetPosts(c.Context())
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Error Fetching Posts")
	}
	return c.Status(fiber.StatusOK).JSON(posts)
}

func HandleCreatePost(c *fiber.Ctx) error {
	post := data.CreatePostRequest{}
	err := c.BodyParser(&post)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Invalid Request Body")
	}

	createdPost, err := data.CreatePost(c.Context(), post)

	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Error Creating Post")
	}

	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func HandleGetPostByID(c *fiber.Ctx) error {
	id, err := ParseID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, err, "Invalid post ID format")
	}

	post := &data.Post{}

	post, err = data.GetPostByID(c.Context(), id)
	if err != nil {
		if err == data.ErrPostNotFound {
			return handleError(c, fiber.StatusNotFound, err, "Post not found")
		}
		return handleError(c, fiber.StatusInternalServerError, err, "Error retrieving post")
	}

	return c.Status(fiber.StatusOK).JSON(post)
}

func HandleUpdatePost(c *fiber.Ctx) error {
	id, err := ParseID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, err, "Invalid post ID format")
	}

	post := data.UpdatePostRequest{}
	err = c.BodyParser(&post)
	if err != nil {
		return handleError(c, fiber.StatusInternalServerError, err, "Invalid Request Body")
	}

	updatedPost, err := data.UpdatePost(c.Context(), id, post)
	if err != nil {
		if err == data.ErrPostNotFound {
			return handleError(c, fiber.StatusNotFound, err, "Post not found")
		}
		return handleError(c, fiber.StatusInternalServerError, err, "Error updating post")
	}

	return c.Status(fiber.StatusOK).JSON(updatedPost)
}

func HandleDeletePost(c *fiber.Ctx) error {
	id, err := ParseID(c)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, err, "Invalid post ID format")
	}

	err = data.DeletePost(c.Context(), id)
	if err != nil {
		if err == data.ErrPostNotFound {
			return handleError(c, fiber.StatusNotFound, err, "Post not found")
		}
		return handleError(c, fiber.StatusInternalServerError, err, "Error deleting Post")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Post deleted successfully",
		"postId":  id,
	})
}
