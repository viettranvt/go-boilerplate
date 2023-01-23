package scripts

import (
	"log"
	"os"

	configs "parishioner_management/internal/configs"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	log.Println("\n--------------------\nLoad env file")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	dbUrl := os.Getenv(configs.EnvMongoDBUrl)

	if dbUrl == "" {
		log.Fatalln("Cannot get DB_URL")
	}
}
