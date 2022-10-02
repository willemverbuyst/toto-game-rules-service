package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")

	return fmt.Sprintf("mongodb://%s:%s@localhost:27017/?authSource=admin", user, password)
}
