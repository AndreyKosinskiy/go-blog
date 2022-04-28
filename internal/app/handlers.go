package app

import (
	"net/http"

	"github.com/AndreyKosinskiy/go-blog/internal/user/model"
	"github.com/AndreyKosinskiy/go-blog/internal/user/repository"
	"github.com/labstack/echo/v4"
)

func validContentType(c echo.Context) bool {
	ct := c.Request().Header.Get("Content-Type")
	if ct == echo.MIMEApplicationXML || ct == echo.MIMEApplicationJSON {
		return true
	}
	return false
}

func (a *App) CreateUserHandle(c echo.Context) error {
	if !validContentType(c) {
		return c.JSON(http.StatusUnsupportedMediaType, "")
	}

	u := &model.User{}
	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	//temp user
	tu, err := model.CreateUser(u.UserName, u.Email, u.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, u)
	}
	r := repository.NewUserRepository(a.Database, a.Logger)
	nu, err := r.Create(c.Request().Context(), tu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, u)
	}
	return c.JSON(http.StatusOK, nu)
}
