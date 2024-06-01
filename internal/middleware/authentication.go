package middleware

import (
	"github.com/Jonatas00/auth_go/internal/authentication"
	"github.com/gin-gonic/gin"
)

func RequireAuthentication(ctx *gin.Context) {
	token, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(401, gin.H{
			"erro": err.Error(),
		})
		return
	}
	if err := authentication.ValidateToken(token); err != nil {
		ctx.JSON(401, gin.H{
			"erro": err.Error(),
		})
	}

	ctx.Next()
}
