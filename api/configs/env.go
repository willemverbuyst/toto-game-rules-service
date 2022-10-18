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

	db := os.Getenv("MONGO_INITDB_DATABASE")

	return fmt.Sprintf("mongodb://localhost:27017/%s", db)

}

func EnvMongoDB() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := os.Getenv("MONGO_INITDB_DATABASE")

	return db
}
