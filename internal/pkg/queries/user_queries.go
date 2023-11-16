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

func (u *User) GetUsers() ([]models.User, error) {
	query := `select username_ from users_`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.Username); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *User) GetUserById(id int) (models.User, error) {
	query := `select username_ from users_ where id_ = $1`

	res := database.DB.QueryRow(query, id)

	var user models.User

	if err := res.Scan(&user.Username); err != nil {
		return user, err
	}

	return user, nil
}
