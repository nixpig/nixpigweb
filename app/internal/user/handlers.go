package user

import (
	"fmt"
	"strconv"

	mp "github.com/geraldo-labs/merge-struct"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Handler = func(c *fiber.Ctx) error

func GetAllHandler(queries UserQueries) Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("before database")

		users, err := queries.GetAll()
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "no users found",
				"count":   0,
				"users":   nil,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "found users",
			"count":   len(users),
			"users":   users,
		})
	}
}

func GetOneHandler(queries UserQueries) Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
				"count":   0,
				"user":    nil,
			})
		}

		user, err := queries.GetOne(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "user with the provided id was not found",
				"user":    nil,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "found user",
			"user":    user,
		})
	}
}

func CreateHandler(queries UserQueries) Handler {
	return func(c *fiber.Ctx) error {
		user := &NewUserModel{}

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		fmt.Println("user:", user)

		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		if err := queries.Create(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "user created",
			"user": NewUserModel{
				Username: user.Username,
				Email:    user.Email,
			},
		})
	}
}

func DeleteHandler(queries UserQueries) Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		user, err := queries.GetOne(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		if err := queries.Delete(user.Id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func UpdateHandler(queries UserQueries) Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		userUpdates := &UserModel{}

		if err := c.BodyParser(userUpdates); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		user, err := queries.GetOne(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		patched, err := mp.Struct(&user, userUpdates)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		if !patched {
			return c.Status(fiber.StatusNotModified).JSON(fiber.Map{
				"error":   false,
				"message": "no changes to make",
				"user":    user,
			})
		}

		if err := queries.Update(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "updated user",
			"user":    user,
		})
	}
}
