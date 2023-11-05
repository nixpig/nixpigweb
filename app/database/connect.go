package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/config"
	"github.com/nixpig/nixpigweb/api/queries"
)

type Queries struct {
	*queries.UserQueries
	*queries.PostQueries
	*queries.ConfigQueries
	*queries.MetaQueries
	*queries.StatusQueries
	*queries.CategoryQueries
}

func Connect() *Queries {
	var err error

	host := config.Get("DBHOST")
	user := config.Get("DBUSER")
	password := config.Get("DBPASSWORD")
	database := config.Get("DBNAME")

	port, err := strconv.Atoi(config.Get("DBPORT"))
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

	return &Queries{
		UserQueries:     &queries.UserQueries{DB: db},
		PostQueries:     &queries.PostQueries{DB: db},
		ConfigQueries:   &queries.ConfigQueries{DB: db},
		MetaQueries:     &queries.MetaQueries{DB: db},
		StatusQueries:   &queries.StatusQueries{DB: db},
		CategoryQueries: &queries.CategoryQueries{DB: db},
	}
}
