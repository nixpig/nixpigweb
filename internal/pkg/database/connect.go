package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
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

func Connect() error {
	var err error

	environment := loadEnvironment()
	connectionString := buildConnectionString(environment)

	wait, err := strconv.Atoi(config.Get("WAIT"))
	if err != nil {
		wait = 5000000000
	}
	time.Sleep(time.Duration(wait))

	fmt.Println("Trying to connect to database: ", connectionString)

	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(fmt.Errorf("error connecting to database\n%v", err))
		return err
	}

	if err = DB.Ping(); err != nil {
		fmt.Println(fmt.Errorf("failed to ping database\n%v", err))
		return err
	}

	fmt.Println("successfully pinged database")

	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get driver from instance\n%v", err))
	}

	m, err := migrate.NewWithDatabaseInstance("github://nixpig/nixpigweb/db/migrations", "postgres", driver)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to create migration\n%v", err))
	}

	m.Up()

	return nil
}

func loadEnvironment() *databaseEnvironment {
	host := config.Get("DATABASE_HOST")
	port := config.Get("DATABASE_PORT")
	name := config.Get("POSTGRES_DB")
	user := config.Get("POSTGRES_USER")
	password := config.Get("POSTGRES_PASSWORD")

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
