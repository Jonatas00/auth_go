package controller

import (
	"github.com/Jonatas00/auth_go/internal/database"
	"github.com/Jonatas00/auth_go/internal/model"
	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user model.User

	ctx.Bind(&user)

	// Check null fields
	if err := user.BeforeCreate(database.DB); err != nil {
		ctx.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(400, result.Error.Error())

		return
	}

	ctx.JSON(200, gin.H{
		"user": &user,
	})
}
