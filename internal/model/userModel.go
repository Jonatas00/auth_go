package model

import (
	"errors"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"notnull;type:varchar(20);default:null"`
	Email    string `json:"email" gorm:"unique;notnull;type:varchar(100);default:null"`
	Password string `json:"password,omitempty" gorm:"notnull;type:varchar(100);default:null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Name == "" {
		return errors.New("name cannot empty")
	}

	if u.Email == "" {
		return errors.New("email cannot empty")
	}
	err := checkmail.ValidateFormat(u.Email)
	if err != nil {
		return errors.New("invalid email")
	}

	if u.Password == "" {
		return errors.New("password cannot empty")
	}

	return nil
}
