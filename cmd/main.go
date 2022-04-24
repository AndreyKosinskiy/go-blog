package main

import (
	"fmt"

	"github.com/AndreyKosinskiy/go-blog/configs"
)

func main() {
	c := configs.New()
	fmt.Println(c)
}
