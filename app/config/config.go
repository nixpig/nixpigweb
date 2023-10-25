package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	return os.Getenv(key)
}
