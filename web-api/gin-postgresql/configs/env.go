package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetDBURI() string {

	LoadEnv()

	return os.Getenv("DB_URI")
}

func DB() string {

	LoadEnv()

	return os.Getenv("DB")

}
