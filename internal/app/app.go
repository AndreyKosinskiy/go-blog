package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AndreyKosinskiy/go-blog/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	Server   http.Server
	Database *gorm.DB
	Logger   *logrus.Logger
}

func New(config *configs.Config) *App {
	s := NewServer(config)
	db := NewORM(config)
	l := NewLogger(config)
	return &App{s, db, l}
}

func (a *App) Run() {
	a.Logger.Info("Server runing...")
	a.setMiddlewares()
	a.setRouters()
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
	a.Logger.Info("Close database connection ...")
	db, err := a.Database.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	a.Logger.Info("Close database connection success!")
	if err := a.Server.Shutdown(ctx); err != nil {
		a.Logger.Fatal(err)
	}
	a.Logger.Info("Gracefull shutdown End.")
}
