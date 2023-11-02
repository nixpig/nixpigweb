package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/nixpig/nixpigweb/api/database"
)

func ValidateUserToken(token *jwt.Token, id int) bool {
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["id"].(float64))

	return id == uid
}

func ValidateRoleToken(token *jwt.Token, role string) bool {
	claims := token.Claims.(jwt.MapClaims)
	tokenRole := claims["role"]

	return tokenRole == role
}

func ValidateAdminToken(token *jwt.Token) bool {
	claims := token.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))
	claimedRole := claims["role"]

	db := database.Connect()

	role, err := db.GetUserRoleById(id)
	if err != nil {
		return false
	}

	return claimedRole == "admin" && claimedRole == role
}
