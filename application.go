package main

import (
	"log"

	"github.com/apmath-web/currency/Application/v1/routing"
)

func main() {
	router := routing.GenRouter()
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
