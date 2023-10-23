package connections

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "lol"
	database = "nixpigweb"
)

func Postgres() (*sql.DB, error) {
	var err error

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s database=%s sslmode=disable",
			host, port, user, password, database,
		),
	)
	if err != nil {
		return db, err
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	return db, nil
}
