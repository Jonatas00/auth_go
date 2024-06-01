package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_PORT  string
	JWT_TOKEN string

	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	if APP_PORT = os.Getenv("APP_PORT"); APP_PORT == "" {
		APP_PORT = "8080"
	}
	JWT_TOKEN = os.Getenv("JWT_TOKEN")

	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
}

func LoadConfigs() {
	loadEnv()
}
