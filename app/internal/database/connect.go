package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/config"
	// "github.com/nixpig/nixpigweb/api/internal/user"
)

// type Queries struct {
// 	*user.UserQueries
// }

func Connect() *sql.DB {
	var err error

	host := config.Config("DBHOST")
	user := config.Config("DBUSER")
	password := config.Config("DBPASSWORD")
	database := config.Config("DBNAME")

	port, err := strconv.Atoi(config.Config("DBPORT"))
	if err != nil {
		panic("failed to get database port from environment")
	}

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		host, port, user, password, database,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("failed to open database connection")
	}

	if err = db.Ping(); err != nil {
		defer db.Close()
		panic("failed to ping database")
	}

	// return &Queries{
	// 	UserQueries: &user.UserQueries{DB: db},
	// }
	return db
}
