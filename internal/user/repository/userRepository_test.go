package repository_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/AndreyKosinskiy/go-blog/internal/user/model"
	"github.com/AndreyKosinskiy/go-blog/internal/user/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func testUser() *model.User {
	userName := "TestUser"
	email := "test@test.user"
	password := "8estpassword"
	tu, err := model.CreateUser(userName, email, password)
	if err != nil {
		log.Fatal(err)
	}
	return tu
}

func newMock() (*gorm.DB, func(), sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}

	orm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	teardown := func() {
		db, _ := orm.DB()
		db.Close()
	}
	return orm, teardown, mock
}

func TestCreateUser(t *testing.T) {
	u := testUser()
	orm, teardown, mock := newMock()
	defer teardown()

	rows := sqlmock.NewRows([]string{"id", "user_name", "email", "password", "is_deleted", "created_at", "update_at", "deleted_at"}).
		AddRow(u.ID, u.UserName, u.Email, u.Password, u.IsDeleted, u.CreatedAt, u.UpdatedAt, u.DeletedAt)
	mock.ExpectBegin()
	mock.ExpectQuery(repository.CreateUserSqlQuery).
		WithArgs(u.ID, u.UserName, u.Email, u.Password, u.IsDeleted, u.CreatedAt, u.UpdatedAt, u.DeletedAt).
		WillReturnRows(rows)
	mock.ExpectCommit()

	r := repository.NewUserRepository(orm, logrus.New())
	nu, err := r.Create(context.Background(), u)

	assert.NoError(t, err)
	assert.NotNil(t, nu)
}

func TestCreateUserEmailError(t *testing.T) {
	t.Skip()
}
func TestDeleteUser(t *testing.T) {
	u := testUser()
	orm, teardown, mock := newMock()
	defer teardown()

	rows := sqlmock.NewRows([]string{"id", "user_name", "email", "password", "is_deleted", "created_at", "update_at", "deleted_at"}).
		AddRow(u.ID, u.UserName, u.Email, u.Password, u.IsDeleted, u.CreatedAt, u.UpdatedAt, u.DeletedAt)
	mock.ExpectBegin()
	mock.ExpectQuery(repository.DeleteUserSqlQuery).WithArgs(true, u.ID, false).WillReturnRows(rows)
	mock.ExpectCommit()

	r := repository.NewUserRepository(orm, logrus.New())
	// du - deleted user
	du, err := r.Delete(context.Background(), u.ID)

	assert.NoError(t, err)
	assert.Equal(t, u.ID, du.ID)
}
func TestDeleteUserError(t *testing.T) {
	id := uuid.New()
	orm, teardown, mock := newMock()
	defer teardown()

	mock.ExpectBegin()
	mock.ExpectQuery(repository.DeleteUserSqlQuery).WithArgs(true, id, false).WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectRollback()

	r := repository.NewUserRepository(orm, logrus.New())
	u, err := r.Delete(context.Background(), id)

	assert.Equal(t, gorm.ErrRecordNotFound, err)
	fmt.Println(u)
	assert.Equal(t, u, &model.User{})
}
