package app

import "github.com/labstack/echo/v4"

func (a *App) setRouters() {
	e := a.Server.Handler.(*echo.Echo)

	//e.GET("/users", a.createUser)
	e.POST("/api/v1/users", a.CreateUserHandle)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
}
func (a *App) setMiddlewares() {

}
