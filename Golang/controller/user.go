package controller

import (
	"mading/helper"
	"mading/model/web"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetUserWithTrx(ctx *gin.Context) {
	req := new(web.UserRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// TODO: check validation here
	if result := helper.TimeCheck(req.Time); !result {
		ctx.JSON(400, gin.H{
			"message": "please wait 1 minute from last request",
		})
		return
	}

	// TODO: call service func
	res, err := c.service.GetUserWithTrx(req.Email, req.Username)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": res,
	})
}
