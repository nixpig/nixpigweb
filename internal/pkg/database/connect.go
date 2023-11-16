package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
)

type databaseEnvironment struct {
	host     string
	port     int
	name     string
	user     string
	password string
}

func Connect() *sql.DB {
	var err error

	environment := loadEnvironment()
	connectionString := buildConnectionString(environment)

	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(fmt.Errorf("error connecting to database\n%v", err))
		os.Exit(1)
	}

	if err = DB.Ping(); err != nil {
		fmt.Println(fmt.Errorf("failed to ping database\n%v", err))
		os.Exit(1)
	}

	return DB
}

func loadEnvironment() *databaseEnvironment {
	host := config.Get("DATABASE_HOST")

	user := config.Get("DATABASE_USER")

	name := config.Get("DATABASE_DB")

	password := config.Get("DATABASE_PASSWORD")

	port := config.Get("DATABASE_PORT")

	portNumber, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(fmt.Errorf("error converting port string to number\n%v", err))
	}

	return &databaseEnvironment{
		host:     host,
		port:     portNumber,
		name:     name,
		user:     user,
		password: password,
	}
}

func buildConnectionString(environment *databaseEnvironment) string {
	return fmt.Sprintf(
		"host=%s port=%d database=%s user=%s password=%s sslmode=disable",
		environment.host, environment.port, environment.name, environment.user, environment.password,
	)
}
