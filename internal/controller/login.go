package controller

import (
	"github.com/Jonatas00/auth_go/internal/database"
	"github.com/Jonatas00/auth_go/internal/middleware"
	"github.com/Jonatas00/auth_go/internal/model"
	"github.com/gin-gonic/gin"
)

type loginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var loginParams loginParams
	ctx.BindJSON(&loginParams)

	var userDB model.User
	database.DB.Select("id", "name", "email", "password").Where("email = ?", loginParams.Email).Find(&userDB)

	err := middleware.CheckPassword(userDB.Password, loginParams.Password)
	if err != nil {
		ctx.JSON(401, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		//"loginParams": loginParams,
		//"user in db":  userDB,
		"message": "Login Successful",
	})
}
