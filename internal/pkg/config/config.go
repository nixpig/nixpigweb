package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const envPath = ".env"

func Init() error {
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println(fmt.Errorf("error loading environment variables\n this might be fatal.\n%v", err))
		return err
	}

	return nil
}

func Get(key string) string {
	return os.Getenv(key)
}
