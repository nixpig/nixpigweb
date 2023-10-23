package queries

import (
	"database/sql"
	"fmt"

	"github.com/nixpig/nixpigweb/api/models"
)

type UserQueries struct {
	Db *sql.DB
}

func (q *UserQueries) GetUsers() ([]models.User, error) {
	fmt.Println("executing db query...")
	users := []models.User{}

	query := "select id, username, email, is_admin, registered_at from users order by id"

	rows, err := q.Db.Query(query)
	fmt.Println("executed query and got rows or error")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	fmt.Println("closed the rows")

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

	query := "select id, username, email, is_admin, registered_at from users where id=$1 limit 1"

	row := q.Db.QueryRow(query, id)

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.RegisteredAt); err != nil {
		return user, err
	}

	return user, nil
}
