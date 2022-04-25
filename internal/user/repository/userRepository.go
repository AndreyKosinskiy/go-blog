package userRepository

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func (r *UserRepository) Create() {}
