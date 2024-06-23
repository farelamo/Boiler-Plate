package controller

import (
	"mading/helper/logger"
	"mading/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *service.Service
	log     *logger.Log
}

func New(svc *service.Service) *Controller {
	return &Controller{
		service: svc,
		log:     logger.New("Controller"),
	}
}

func (c *Controller) RegisterRoute(app *gin.Engine) {
	group := app.Group("/api/v1")

	group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	
	// TODO: call service func
	group.GET("/users", c.GetUserWithTrx)
}
