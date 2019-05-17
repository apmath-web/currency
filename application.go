package main

import (
	"github.com/apmath-web/currency/Application/v1"
	"log"
)

func main() {
	router := v1.GenRouter()
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
