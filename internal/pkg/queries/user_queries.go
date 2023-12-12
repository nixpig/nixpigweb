package queries

import (
	"github.com/nixpig/nixpigweb/internal/pkg/database"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
)

func CreateUser(user *models.User) (int64, error) {
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

func GetUsers() ([]models.User, error) {
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

func GetUserById(id int) (models.User, error) {
	query := `select id_, username_, email_, is_admin_ from users_ where id_ = $1`

	res := database.DB.QueryRow(query, id)

	var user models.User

	if err := res.Scan(&user.Id, &user.Username, &user.Email, &user.IsAdmin); err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(user *models.User) (int64, error) {
	query := `update users_ set username_ = $2, email_ = $3, password_ = $4 where id_ = $1`

	res, err := database.DB.Exec(query, &user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
