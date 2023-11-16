package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

func ValidateUserToken(token *jwt.Token, id int) bool {
	claims := token.Claims.(jwt.MapClaims)
	claim_id := int(claims["user_id"].(float64))

	return claim_id == id

}

func ValidateAdminToken(token *jwt.Token, isAdmin bool) bool {
	claims := token.Claims.(jwt.MapClaims)
	claim_is_admin := claims["is_admin"]

	return claim_is_admin == isAdmin
}
