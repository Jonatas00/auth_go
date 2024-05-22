package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8008"
	}
}

func Load() {
	go loadEnv()
}
