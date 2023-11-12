package main

import (
	"fmt"

	"github.com/nixpig/nixpigweb/internal/api/server"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
)

func main() {
	config.Init()

	server.Start()
	// db := database.Connect()
	//
	// contentQueries := queries.Content{DB: db}
	//
	// _, err := contentQueries.CreateContent(&models.Content{
	// 	Title:    "Some title",
	// 	Subtitle: "Subtitle",
	// 	Slug:     "some-slug",
	// 	Body:     "Some body content",
	// 	Type:     "post",
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("error inserting in main\n%v", err))
	// }
	//
	// content, err := contentQueries.GetContent()
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("error getting in main\n%v", err))
	// } else {
	// 	fmt.Println(fmt.Printf("success getting in main\n%v", content))
	// }

	fmt.Println("Hello, api!")
}
