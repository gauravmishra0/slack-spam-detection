package controllers

import (
	"SlackNotification/api/models"
	"SlackNotification/api/services/slacksvc"
	"SlackNotification/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Slack(c *gin.Context) {
	log := utils.GetLogCtx(c)
	slack := slacksvc.NewSlackService()

	requestBody := models.SlackArgs{}
	if err := c.BindJSON(&requestBody); err != nil {
		log.Errorf("Invalid payload format %v", err)
		c.JSON(http.StatusBadRequest, models.CommonAPIResponse{
			IsError:         true,
			ResponseMessage: "Invalid payload format",
			ResponseData:    err,
		})
		return
	}

	if requestBody.Type == "SpamNotification" {
		msg, err := slack.PostMessageToChannel(c, requestBody.Email)
		if err != nil {
			logrus.Errorf("Unable to send message %v", err)
			c.JSON(http.StatusBadRequest, models.CommonAPIResponse{
				IsError:         true,
				ResponseMessage: "Unable to send message",
				ResponseData:    err,
			})
			return
		}

		logrus.Infof("Message sent %s", msg)
		c.JSON(http.StatusOK, models.CommonAPIResponse{
			IsError:         false,
			ResponseMessage: "Message sent",
			ResponseData:    msg,
		})

		return
	}

	logrus.Infof("Not a spam message")
	c.JSON(http.StatusBadRequest, models.CommonAPIResponse{
		IsError:         false,
		ResponseMessage: "Not a spam message",
		ResponseData:    "",
	})
}
