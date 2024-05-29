package controller

import (
	"github.com/Jonatas00/auth_go/internal/database"
	"github.com/Jonatas00/auth_go/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(loginParams.Password))
	if err != nil {
		ctx.JSON(401, gin.H{
			"erro": "invalid password",
		})
		return
	}

	ctx.JSON(200, gin.H{
		//"loginParams": loginParams,
		//"user in db":  userDB,
		"message": "Login Successful",
	})
}
