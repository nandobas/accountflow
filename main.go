package main

import (
	"accountflow/api"
	"accountflow/environment"
	"fmt"
)

func init() {
	environment.InitEnv(".env")
}

func main() {
	fmt.Println("init account flow")

	api.NewService().Start()
}
