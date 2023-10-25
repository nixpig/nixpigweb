package user

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type UserQueries struct {
	*sql.DB
}

func (q *UserQueries) GetAll() ([]UserModel, error) {
	users := []UserModel{}

	query := "select id, username, email, is_admin, registered_at from users order by id"

	rows, err := q.Query(query)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		user := UserModel{}

		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.RegisteredAt); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (q *UserQueries) GetOne(id int) (UserModel, error) {
	user := UserModel{}

	query := "select id, username, email, is_admin, registered_at from users where id = $1 limit 1"

	row := q.QueryRow(query, id)

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.RegisteredAt); err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) Create(user *NewUserModel) error {
	query := "insert into users (username, email, is_admin, password, registered_at) values ($1, $2, $3, $4, $5)"

	_, err := q.Exec(query, &user.Username, &user.Email, false, &user.Password, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (q *UserQueries) Delete(id int) error {
	query := "delete from users where id = $1"

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (q *UserQueries) Update(user *UserModel) error {
	query := "update users set username = $2, email = $3, is_admin = $4, password = $5 where id = $1"

	_, err := q.Exec(query, &user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.Password)
	if err != nil {
		return err
	}

	return nil
}
