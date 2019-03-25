package main

import (
	"github.com/apmath-web/currency/Application/routing"
	"github.com/apmath-web/currency/Infrastructure/applicationModels"
	"github.com/apmath-web/currency/Infrastructure/repositories"
	"log"
)

func main() {
	repositories.Repo.PutModel(applicationModels.GenHelloWorldApplicationModel("hello_world"))
	router := routing.GenRouter()
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
