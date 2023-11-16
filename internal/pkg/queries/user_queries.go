package queries

import (
	"github.com/nixpig/nixpigweb/internal/pkg/database"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
)

type User struct{}

func (u *User) CreateUser(user *models.User) (int64, error) {
	query := `insert into users_ (username_, email_, password_) values ($1, $2, $3)`

	res, err := database.DB.Exec(query, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
