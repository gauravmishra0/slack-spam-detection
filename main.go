package main

import (
	"SlackNotification/api"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	api.Run()
}
