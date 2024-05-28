package middleware

import (
	"errors"

	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(password string) (string, error) {
	const minEntropyBits = 15
	err := passwordvalidator.Validate(password, minEntropyBits)
	if err != nil {
		return "", errors.New("weak password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
