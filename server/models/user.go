package models

import (
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username      string `json:"username" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	PasswordCheck string `json:"passwordCheck" binding:"required,eqfield=Password"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (b *RegisterBody) SetNewUser() *NewUser {
	// TODO handle error
	bytes, _ := bcrypt.GenerateFromPassword([]byte(b.Password), 14)
	return &NewUser{
		Username: b.Username,
		Email:    b.Email,
		Password: string(bytes),
	}
}
