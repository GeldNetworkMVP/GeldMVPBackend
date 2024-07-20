package configs

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	port    = ""
	EnvName = ""
)

func LoadEnv() {
	godotenv.Load(".env")
	EnvName = os.Getenv("BRANCH_NAME")
	port = os.Getenv("BE_PORT")
}

func GetPort() string {
	LoadEnv()
	if port != "" {
		return ":" + port
	}
	return ":8000"
}
