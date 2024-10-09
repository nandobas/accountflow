package main

import (
	"accountflow/api"
	"accountflow/environment"

	"github.com/sirupsen/logrus"
)

func init() {
	environment.InitEnv(".env")
}

func main() {
	logrus.Info("init account flow")

	api.NewService().Start()
}
