package queries

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/models"
)

type UserQueries struct {
	*sql.DB
}

func (q *UserQueries) GetUsers() ([]models.User, error) {
	users := []models.User{}

	query := "select id, username, email, is_admin, registered_at from users order by id"

	rows, err := q.Query(query)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		user := models.User{}

		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.RegisteredAt); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (q *UserQueries) GetUser(id int) (models.User, error) {
	user := models.User{}

	query := "select id, username, email, is_admin, registered_at from users where id = $1 limit 1"

	row := q.QueryRow(query, id)

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.RegisteredAt); err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) CreateUser(user *models.NewUser) error {
	query := "insert into users (username, email, is_admin, password, registered_at) values ($1, $2, $3, $4, $5)"

	_, err := q.Exec(query, &user.Username, &user.Email, false, &user.Password, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (q *UserQueries) DeleteUser(id int) error {
	query := "delete from users where id = $1"

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// TODO: update user
