package main

import (
	"log"

	"github.com/Jonatas00/auth_go/internal/config"
	"github.com/Jonatas00/auth_go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	r := gin.Default()
	routes.LoadRoutes(&r.RouterGroup)

	if err := r.Run(":" + config.PORT); err != nil {
		log.Fatal(err)
	}
}
