package controllers

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nixpig/nixpigweb/api/config"
	"github.com/nixpig/nixpigweb/api/database"
	"github.com/nixpig/nixpigweb/api/models"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswordHash(password, hash string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func isEmail(test string) bool {
	_, err := mail.ParseAddress(test)
	if err != nil {
		return false
	}

	return true
}

func getUserByEmail(email string) (models.User, error) {
	db := database.Connect()

	user, err := db.GetUserByEmail(email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func getUserById(id int) (models.User, error) {
	db := database.Connect()

	user, err := db.GetUserById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func Login(c *fiber.Ctx) error {
	input := &models.Login{}

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid input",
			"data":    nil,
		})
	}

	user := models.User{}
	var err error

	db := database.Connect()

	if isEmail(input.Username) {
		user, err = db.GetUserByEmail(input.Username)
	} else {
		user, err = db.GetUserByUsername(input.Username)
	}
	if err != nil {
		fmt.Println("get user error:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "invalid login provided",
			"data":    nil,
		})
	}

	fmt.Println("user found and assigned for login:", user)

	isAuthorised := ComparePasswordHash(input.Password, user.Password)
	if !isAuthorised {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "invalid login provided",
			"data":    nil,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = user.Username
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour + 72).Unix()

	signedToken, err := token.SignedString([]byte(config.Get("SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "an error occurred",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "authorised",
		"data":    map[string]string{"token": signedToken},
	})
}
