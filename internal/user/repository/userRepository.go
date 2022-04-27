package repository

import (
	"context"
	"fmt"

	"github.com/AndreyKosinskiy/go-blog/internal/user/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewUserRepository(db *gorm.DB, logger *logrus.Logger) *UserRepository {
	return &UserRepository{db, logger}
}
func (r *UserRepository) Create(ctx context.Context, u *model.User) (*model.User, error) {
	r.db.Table("users").Create(u)
	fmt.Println(u)
	return u, nil
}
