package app

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AndreyKosinskiy/go-blog/configs"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type App struct {
	Server   http.Server
	Database *sql.DB
	Logger   *logrus.Logger
}

func NewServer(config *configs.Config) http.Server {
	e := echo.New()
	e.GET("/hello", hello)
	s := http.Server{
		Addr:    ":" + config.Port,
		Handler: e,
	}
	return s
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func NewDatabase(config *configs.Config) *sql.DB {
	// db, err := sql.Open("postgres", config.DbURL)
	// if err != nil {
	// 	log.Fatal("Can`t open database by URL: ", config.DbURL)
	// }
	return nil
}

func NewLogger(config *configs.Config) *logrus.Logger {
	l := logrus.New()
	return l
}

func New(config *configs.Config) *App {
	s := NewServer(config)
	db := NewDatabase(config)
	l := NewLogger(config)
	return &App{s, db, l}
}

func (a *App) Run() {
	a.Logger.Info("Server runing...")
	go func() {
		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	a.Logger.Info("Gracefull shutdown Begin...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.Server.Shutdown(ctx); err != nil {
		a.Logger.Fatal(err)
	}
	a.Logger.Info("Gracefull shutdown End.")
}
