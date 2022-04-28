package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id" xml:"id"`
	UserName  string    `gorm:"user_name" json:"userName" xml:"userName"`
	Email     string    `json:"email" xml:"email"`
	Password  string    `json:"-" xml:"-"`
	IsDeleted bool      `json:"-" xml:"-"`
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at"`
	DeletedAt time.Time `json:"-" xml:"-"`
}

func CreateUser(userName string, email string, password string) (*User, error) {
	nu := &User{UserName: userName, Email: email, Password: password}
	err := nu.BeforeCreate()
	if err != nil {
		return nil, err
	}
	return nu, nil
}

func (u *User) BeforeCreate() (err error) {
	u.ID = uuid.New()
	u.Password, err = passwordHashing(u.Password)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func passwordHashing(password string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hp), nil
}
