package queries

import (
	"database/sql"

	"github.com/nixpig/nixpigweb/internal/pkg/models"
)

type User struct {
	*sql.DB
}

func (u *User) CreateUser(user *models.User) (int64, error) {
	query := `insert into users_ (username_, email_, password_) values ($1, $2, $3)`

	res, err := u.Exec(query, &user)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
