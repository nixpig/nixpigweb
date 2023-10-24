package queries

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/models"
)

type UserQueries struct {
	DB *sql.DB
}

func (q *UserQueries) GetUsers() ([]models.User, error) {
	users := []models.User{}

	query := "select id, username, email, is_admin, registered_at from users order by id"

	rows, err := q.DB.Query(query)
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

	query := "select id, username, email, is_admin, registered_at from users where id=$1 limit 1"

	fmt.Println("before query row")
	row := q.DB.QueryRow(query, id)
	fmt.Println("after query row")

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.IsAdmin, &user.RegisteredAt); err != nil {
		return user, err
	}

	return user, nil
}
