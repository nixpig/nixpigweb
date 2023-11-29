package main

import (
	"fmt"
	"os"

	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/database"
	"github.com/nixpig/nixpigweb/internal/web/server"
)

func main() {
	if err := config.Init(); err != nil {
		fmt.Println(fmt.Errorf("failed to load .env file\nthis shouldn't be fatal if environment variables are set by some other means\n%v", err))
	}

	if err := database.Connect(); err != nil {
		fmt.Println(fmt.Errorf("failed to connect to database\n%v", err))
		os.Exit(1)
	}

	contextPath := config.Get("WEB_CONTEXT")
	port := config.Get("WEB_PORT")

	server.Start(contextPath, port)
}
