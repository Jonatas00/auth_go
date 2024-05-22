package controller

import (
	"encoding/json"

	"github.com/Jonatas00/auth_go/internal/model"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	reqBody, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	var user model.UserModel
	if err = json.Unmarshal(reqBody, &user); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	if err = user.Prepare(); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, gin.H{
		"user": user,
	})
}
