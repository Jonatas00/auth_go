package model

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type UserModel struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u UserModel) validate() error {
	var invalidFields []string
	if u.Name == "" {
		invalidFields = append(invalidFields, "name")
	}
	if u.Email == "" {
		invalidFields = append(invalidFields, "email")
	}
	if u.Password == "" {
		invalidFields = append(invalidFields, "password")
	}

	if len(invalidFields) > 0 {
		msg := "Invalid fields: "
		for _, field := range invalidFields {
			msg += field + ", "
		}

		return errors.New(msg)
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email format")
	}

	return nil
}

func (u UserModel) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}

	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)

	return nil
}
