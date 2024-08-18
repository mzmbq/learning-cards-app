package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password,omitempty" validate:"required,min=6"`
	EncryptedPassword string `json:"-"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) BeforeCreate() error {
	if u.Password == "" {
		return nil
	}

	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(b)

	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	return err == nil
}
