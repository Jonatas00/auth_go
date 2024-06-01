package controller

import (
	"fmt"
	"net/http"

	"github.com/Jonatas00/auth_go/internal/authentication"
	"github.com/Jonatas00/auth_go/internal/database"
	"github.com/Jonatas00/auth_go/internal/model"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var loginParams model.Login
	ctx.BindJSON(&loginParams)

	var userDB model.User
	database.DB.Select("id", "name", "email", "password").Where("email = ?", loginParams.Email).Find(&userDB)

	err := authentication.CheckPassword(userDB.Password, loginParams.Password)
	if err != nil {
		ctx.JSON(401, gin.H{
			"erro": err.Error(),
		})
		return
	}

	token, err := authentication.GenerateToken(fmt.Sprint(userDB.ID))
	if err != nil {
		ctx.JSON(401, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)

	// Cookie name, value, time in seconds, path, domain, secure, httponly
	ctx.SetCookie("Authorization", token, 3600*24, "", "", false, true)
	ctx.JSON(200, gin.H{
		//"loginParams": loginParams,
		//"user in db":  userDB,
		"message": "Login Successful",
		"token":   token,
	})
}
