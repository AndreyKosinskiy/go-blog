package app

import (
	"context"
	"log"
	"net/http"

	"github.com/AndreyKosinskiy/go-blog/configs"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewServer(config *configs.Config) http.Server {
	e := echo.New()
	s := http.Server{
		Addr:    ":" + config.Port,
		Handler: e,
	}
	return s
}

func NewDatabase(config *configs.Config) *pgxpool.Pool {
	log.Printf("Init Database ...")
	db, err := pgxpool.Connect(context.Background(), config.DbURL)
	if err != nil {
		log.Fatalf("Can`t open database by URL: %s", err)
	}
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("Connection error, can`t ping database: %s", err)
	}
	log.Printf("Init Database success!")
	return db
}

func NewORM(config *configs.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewLogger(config *configs.Config) *logrus.Logger {
	l := logrus.New()
	return l
}
