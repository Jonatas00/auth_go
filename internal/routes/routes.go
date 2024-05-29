package routes

import (
	"github.com/Jonatas00/auth_go/internal/controller"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.RouterGroup) {
	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)

	r.GET("/getusers", controller.Getusers)
}
