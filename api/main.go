package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/controllers"
)

var db *sql.DB

func main() {

	fmt.Println("creating new fiber app")
	app := fiber.New()

	fmt.Println("registering get user route")
	app.Get("/user", controllers.GetUsers)

	log.Fatal(app.Listen(":3000"))
}
