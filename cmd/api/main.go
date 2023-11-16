package main

import (
	"fmt"
	"os"

	"github.com/nixpig/nixpigweb/internal/api/server"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/database"
)

func main() {
	if err := config.Init(); err != nil {
		fmt.Println(fmt.Errorf("error initialising app config\n%v", err))
		os.Exit(1)
	}

	database.Connect()
	server.Start()
}
