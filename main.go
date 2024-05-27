package main

import (
	"log"

	"github.com/Jonatas00/auth_go/internal/config"
	"github.com/Jonatas00/auth_go/internal/database"
	"github.com/Jonatas00/auth_go/internal/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfigs()
	database.Connect()
	database.LoadMigrations()
}

func main() {
	r := gin.Default()

	routes.LoadRoutes(&r.RouterGroup)

	if err := r.Run(":" + config.APP_PORT); err != nil {
		log.Fatal(err)
	}
}
