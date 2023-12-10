package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
	"github.com/nixpig/nixpigweb/internal/pkg/services"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	var input models.Login

	if err := c.BodyParser(&input); err != nil {
		fmt.Println("ERROR: failed to parse login input\n", err)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	user, err := queries.GetUserByUsername(input.Username)
	if err != nil {
		fmt.Println("ERROR: failed to get user by username\n", err)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	if !comparePasswordHash(user.Password, input.Password) {
		fmt.Println("ERROR: password doesn't match")

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	userId := user.Id
	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	claims["user_id"] = userId
	claims["is_admin"] = user.IsAdmin
	claims["exp"] = expiresAt
	claims["iat"] = issuedAt

	signedToken, err := token.SignedString([]byte(config.Get("SECRET")))
	if err != nil {
		fmt.Println("ERROR: failed to get signed token string\n", err)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	queries.DeleteSessionsByUserId(userId)

	session := models.Session{
		UserId:    userId,
		Token:     signedToken,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}

	sessionCount, err := queries.SaveSession(session)
	if err != nil || sessionCount != 1 {
		fmt.Println("ERROR: failed to save user session\n", err)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "authorised",
		"data":    map[string]string{"token": signedToken},
	})

}

func comparePasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func Logout(c *fiber.Ctx) error {
	token := c.Locals("user")

	if token == nil {
		fmt.Println("ERROR: failed to get toke from session")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "internal server error - you may not have been logged out",
			"data":    nil,
		})
	}

	claims := token.(*jwt.Token).Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	session, err := queries.GetSessionByUserId(userId)
	if err != nil {
		fmt.Println("ERROR: failed to get session by user id\n", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{

			"error":   true,
			"message": "internal server error - you may not have been logged out",
			"data":    nil,
		})
	}

	deletedSessionsCount, err := queries.DeleteSessionsBySessionId(session.Id)
	if err != nil || deletedSessionsCount != 1 {
		fmt.Println("ERROR: failed to delete existing sessions\n", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "internal server error - you may not have been logged out",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "you have been logged out",
		"data":    nil,
	})
}

func ChangePassword(c *fiber.Ctx) error {
	var changePassword models.ChangePassword

	token := c.Locals("user")
	if token == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	if !services.ValidateUserToken(token.(*jwt.Token)) {
		fmt.Println("ERROR: failed to validate user token")

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	if err := c.BodyParser(&changePassword); err != nil {
		fmt.Println("ERROR: failed to parse password change data from request")

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	user, err := queries.GetUserByUsername(changePassword.Username)
	if err != nil {
		fmt.Println("ERROR: failed to get user by username\n", err)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	if err != nil {
		fmt.Println("ERROR: failed to hash old password")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "internal server error - password not changed",
			"data":    nil,
		})
	}

	if !comparePasswordHash(user.Password, string(changePassword.OldPassword)) {
		fmt.Println("ERROR: password doesn't match")

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(changePassword.NewPassword), 14)

	if err != nil {
		fmt.Println("ERROR: failed to hash new password")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "internal server error - password not changed",
			"data":    nil,
		})
	}

	passwordChanged, err := queries.ChangePassword(changePassword.Username, string(hashedNewPassword))
	if err != nil || !passwordChanged {
		fmt.Println("ERROR: failed to execute query to change password\n", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "internal server error - password not changed",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "password changed",
		"data":    nil,
	})
}
