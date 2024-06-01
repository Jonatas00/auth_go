package authentication

import (
	"github.com/Jonatas00/auth_go/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = config.JWT_TOKEN

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JWT_KEY))
}
