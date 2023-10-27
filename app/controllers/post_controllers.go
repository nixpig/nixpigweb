package controllers

import (
	"fmt"
	"strconv"

	mp "github.com/geraldo-labs/merge-struct"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"

	"github.com/nixpig/nixpigweb/api/database"
	"github.com/nixpig/nixpigweb/api/models"
)

func validateUserToken(token *jwt.Token, id int) bool {
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["id"].(float64))

	return id == uid
}

func GetPosts(c *fiber.Ctx) error {
	db := database.Connect()

	posts, err := db.GetPosts()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "no posts found",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("found %v posts", len(posts)),
		"data":    posts,
	})
}

func GetPost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse user id",
			"data":    nil,
		})
	}

	db := database.Connect()

	post, err := db.GetPost(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "post with the provided ID was not found",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   true,
		"message": "found post",
		"data":    post,
	})
}

func CreatePost(c *fiber.Ctx) error {

	post := &models.NewPost{}

	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse post data from request body",
			"data":    nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(post); err != nil {
		fmt.Println("failed validation")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to validate post data",
			"data":    nil,
		})
	}

	fmt.Println("locals:", c.Locals("user").(*jwt.Token))
	token := c.Locals("user").(*jwt.Token)

	isValidUserToken := validateUserToken(token, post.UserId)

	if !isValidUserToken {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "user id for post doesn't match authenticated user token",
			"data":    nil,
		})
	}

	db := database.Connect()
	if err := db.CreatePost(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to create new post",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "post published",
		"data": models.NewPost{
			Title: post.Title,
			Body:  post.Body,
		},
	})
}

func DeletePost(c *fiber.Ctx) error {
	// TODO: ensure that user owns the post or is admin
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse user id",
			"data":    nil,
		})
	}

	db := database.Connect()

	post, err := db.GetPost(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "post not found",
			"data":    nil,
		})
	}

	if err := db.DeletePost(post.Id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to delete post from database",
			"data":    nil,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdatePost(c *fiber.Ctx) error {
	// TODO: ensure that user owns post or is admin
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse user id",
			"data":    nil,
		})
	}

	fmt.Println("locals:", c.Locals("user").(*jwt.Token))

	db := database.Connect()
	post, err := db.GetPost(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to find post",
			"data":    nil,
		})
	}

	postUpdates := &models.Post{}

	if err := c.BodyParser(postUpdates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse post update from request body",
			"data":    nil,
		})
	}

	patched, err := mp.Struct(&post, postUpdates)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to merge post changes",
			"data":    nil,
		})
	}

	if !patched {
		return c.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error":   false,
			"message": "no changes made",
			"data":    post,
		})
	}

	if err := db.UpdatePost(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to save updated post to database",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "updated post",
		"data":    post,
	})
}
