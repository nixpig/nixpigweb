package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
)

func ValidateUserToken(token *jwt.Token, userId int) bool {
	claims := token.Claims.(jwt.MapClaims)
	claim_id := int(claims["user_id"].(float64))
	claim_exp := int(claims["exp"].(float64))

	session, err := queries.GetSessionByToken(token.Raw)
	if err != nil {
		fmt.Println("Error: failed to get session by token\n", err)

		return false
	}

	claimUserMatchesUser := claim_id == userId
	claimUserMatchesSessionUser := claim_id == session.UserId
	tokenIsNotExpired := int64(claim_exp) >= time.Now().Unix()

	return claimUserMatchesUser && claimUserMatchesSessionUser && tokenIsNotExpired
}

func ValidateAdminToken(token *jwt.Token, isAdmin bool) bool {
	claims := token.Claims.(jwt.MapClaims)
	claim_is_admin := claims["is_admin"]

	return claim_is_admin == isAdmin
}
