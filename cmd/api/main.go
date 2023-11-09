package main

import (
	"fmt"

	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/database"
)

func main() {
	config.Init()
	database.Connect()
	fmt.Println("Hello, api!")
}
