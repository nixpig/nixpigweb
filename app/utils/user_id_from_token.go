package utils

import "github.com/golang-jwt/jwt/v5"

func UserIdFromToken(token *jwt.Token) int {
	claims := token.Claims.(jwt.MapClaims)

	return int(claims["id"].(float64))
}
