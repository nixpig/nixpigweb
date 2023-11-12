package main

import (
	"github.com/nixpig/nixpigweb/internal/api/server"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
)

func main() {
	config.Init()
	server.Start()
}
