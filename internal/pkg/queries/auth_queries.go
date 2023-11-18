package queries

import (
	"github.com/nixpig/nixpigweb/internal/pkg/database"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
)

func GetUserByUsername(username string) (models.User, error) {
	query := `select id_, username_, email_, password_, is_admin_ from users_ where username_ = $1`

	var user models.User

	row := database.DB.QueryRow(query, username)
	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.IsAdmin); err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	query := `select id_, username_, email_, password_, is_admin_ from users_ where email_ = $1`

	var user models.User

	row := database.DB.QueryRow(query, email)

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.IsAdmin); err != nil {
		return user, err
	}

	return user, nil
}
