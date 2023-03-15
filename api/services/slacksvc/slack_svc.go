package slacksvc

import (
	"SlackNotification/api/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var SlackSvc SlackActivitiesI = NewSlackService()

type SlackActivitiesI interface {
	PostMessageToChannel(*gin.Context, string) (string, error)
}

type SlackActivities struct {
}

func NewSlackService() *SlackActivities {
	slack := new(SlackActivities)
	return slack
}

func (s SlackActivities) PostMessageToChannel(c *gin.Context, msg string) (string, error) {
	log := utils.GetLogCtx(c)

	url := os.Getenv("SLACK_WEBHOOK_URL")
	data := fmt.Sprintf(`{"text":"%s"}`, msg)

	payload := strings.NewReader(data)

	client := &http.Client{
		Transport: &utils.Request{Transport: http.DefaultTransport, Ctx: c},
		Timeout:   2 * time.Second,
	}
	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("Content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		fields := map[string]interface{}{
			"status_code": res.StatusCode,
			"response":    string(body),
		}
		log.WithFields(fields).Errorf("failed to send slack message")

		return "", errors.New("failed to send slack message")
	}

	log.Info("request success")

	return string(body), nil
}
