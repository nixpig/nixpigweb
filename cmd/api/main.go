package main

import (
	"github.com/nixpig/nixpigweb/internal/api/server"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/database"
)

func main() {
	config.Init()
	database.Connect()
	server.Start()
}
