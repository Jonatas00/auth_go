package controller

import (
	"github.com/Jonatas00/auth_go/internal/database"
	"github.com/Jonatas00/auth_go/internal/model"
	"github.com/gin-gonic/gin"
)

func Getusers(ctx *gin.Context) {
	users := []model.User{}
	database.DB.Find(&users)
	ctx.JSON(200, users)
}
