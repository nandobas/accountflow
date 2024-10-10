package main

import (
	"accountflow/api"
	"accountflow/environment"
	"accountflow/modules/system/lcache"

	"github.com/sirupsen/logrus"
)

func init() {
	environment.InitEnv(".env")
}

func main() {
	logrus.Info("init account flow")

	lcache.InitLocalCache()

	api.NewService().Start()
}
