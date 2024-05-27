package database

import (
	"fmt"
	"log"

	"github.com/Jonatas00/auth_go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT,
	)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
