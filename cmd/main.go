package main

import (
	"github.com/AndreyKosinskiy/go-blog/configs"
	"github.com/AndreyKosinskiy/go-blog/internal/app"
)

func main() {
	confg := configs.New()
	app := app.New(confg)
	app.Run()
}
