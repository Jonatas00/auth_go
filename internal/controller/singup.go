package controller

import (
	"errors"

	"github.com/Jonatas00/auth_go/internal/authentication"
	"github.com/Jonatas00/auth_go/internal/database"
	"github.com/Jonatas00/auth_go/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	hashedPassword, err := authentication.EncryptPassword(user.Password)
	if err != nil {
		ctx.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	user.Password = hashedPassword

	// Check duplicate keys in db and create line
	result := database.DB.Create(&user)
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		ctx.JSON(400, gin.H{
			"erro": "This email is already registered",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"user_id": &user.ID,
		"message": "successfully signed up",
	})
}
