package controllers

import (
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			slack := v1.Group("slack")
			{
				slack.POST("/message", Slack)
			}
		}
	}
}
