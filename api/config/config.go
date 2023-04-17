package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_DB       string
)

func LoadConfig() {
	if err := godotenv.Load(".env.dev"); err != nil {
		log.Fatal("failed load env")
	}
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	POSTGRES_DB = os.Getenv("POSTGRES_DB")
}
