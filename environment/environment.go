package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ListenHttpPort string
)

func InitEnv(pathEnv string) {

	if err := godotenv.Load(pathEnv); err != nil {
		log.Fatal(err)
	}
	ListenHttpPort = os.Getenv("MODULEHTTPPORT")
}
