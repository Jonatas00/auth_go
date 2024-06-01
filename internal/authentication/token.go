package authentication

import (
	"fmt"

	"github.com/Jonatas00/auth_go/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"userID":     userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_TOKEN))
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JWT_TOKEN), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["userID"])
	} else {
		return err
	}

	return nil
}
