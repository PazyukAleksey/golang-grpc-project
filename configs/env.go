package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoUrl() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGOURI")
}
