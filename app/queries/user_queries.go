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

func (q *UserQueries) GetUserById(id int) (models.User, error) {
	user := models.User{}

	query := "select id, username, email, is_admin, registered_at from users where id = $1 limit 1"

	row := q.QueryRow(query, id)

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.RegisteredAt); err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	query := "select id, username, email, is_admin, password from users where email = $1"

	row := q.QueryRow(query, email)

	user := models.User{}

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) GetUserByUsername(username string) (models.User, error) {
	query := "select id, username, email, is_admin, password from users where username = $1"

	row := q.QueryRow(query, username)

	user := models.User{}

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.Password); err != nil {
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

func (q *UserQueries) UpdateUser(user *models.User) error {
	query := "update users set username = $2, email = $3, is_admin = $4, password = $5 where id = $1"

	_, err := q.Exec(query, &user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.Password)
	if err != nil {
		return err
	}

	return nil
}
