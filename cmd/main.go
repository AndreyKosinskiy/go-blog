package main

import (
	"github.com/AndreyKosinskiy/go-blog/configs"
	"github.com/AndreyKosinskiy/go-blog/internal/app"
)

func main() {
	config := configs.New()
	app := app.New(config)
	app.Run()
}
