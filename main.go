package main

import (
	"fmt"
	"mading/config"
	"mading/controller"
	"mading/helper/logger"
	"mading/repository"
	"mading/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var log = logger.New("Main")

func main() {
	validate 	:= validator.New()
	db 			:= config.NewDB()
	repository 	:= repository.New(db)
	service 	:= service.New(repository, validate)
	controller 	:= controller.New(service)

	app := gin.Default()

	controller.RegisterRoute(app)

	app.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
}
