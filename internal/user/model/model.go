package model

import (
	"encoding/xml"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	XMLName  xml.Name  `json:"-" xml:"user"`
	Id       uuid.UUID `json:"id" xml:"id"`
	UserName string    `json:"userName" xml:"userName"`
	Email    string    `json:"email" xml:"email"`
	Password string    `json:"password" xml:"password"`
}

func BeforeCreate(u *User) (*User, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hp)
	return u, nil
}
