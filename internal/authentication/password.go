package authentication

import (
	"errors"

	//passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	// const minEntropyBits = 25
	// err := passwordvalidator.Validate(password, minEntropyBits)
	// if err != nil {
	// 	return "", errors.New("weak password")
	// }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(DBPassword, reqPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(DBPassword), []byte(reqPassword))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
