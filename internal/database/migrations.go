package database

import "github.com/Jonatas00/auth_go/internal/model"

func LoadMigrations() {
	DB.AutoMigrate(&model.User{})
}
