package utils

import "github.com/golang-jwt/jwt/v5"

func ValidateUserToken(token *jwt.Token, id int) bool {
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["id"].(float64))

	return id == uid
}

func ValidateAdminToken(token *jwt.Token) bool {
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"]

	return role == "admin"
}
