package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/AndreyKosinskiy/go-blog/internal/user/model"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewUserRepository(db *gorm.DB, logger *logrus.Logger) *UserRepository {
	return &UserRepository{db, logger}
}

func RepoError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		fmt.Printf("\n%#v", pgErr)

		switch pgErr.Code {
		case "23505":
			err = errors.New("value" + pgErr.ColumnName + "must be unique")
		}
	}
	return err
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) (*model.User, error) {
	res := r.db.Clauses(clause.Returning{}).Create(u)
	return u, RepoError(res.Error)
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) (*model.User, error) {
	tu := &model.User{}
	res := r.db.Model(tu).Clauses(clause.Returning{}).
		Where("id = ?", id).
		Where("is_deleted = ?", false).
		UpdateColumn("is_deleted", true)

	if res.RowsAffected == 0 {
		res.Error = gorm.ErrRecordNotFound
		tu = &model.User{}
	}
	return tu, RepoError(res.Error)
}
