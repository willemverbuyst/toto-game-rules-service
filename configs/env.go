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
	db := os.Getenv("MONGO_INITDB_DATABASE")

	return fmt.Sprintf("mongodb://%s:%s@mongodb:27017/%s?authSource=admin", user, password, db)
}

func EnvMongoDB() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := os.Getenv("MONGO_INITDB_DATABASE")

	return db
}
