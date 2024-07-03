package configs

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	port          = ""
	EnvName       = ""
	nftBackendUrl = ""
)

func LoadEnv() {
	godotenv.Load(".env")
	EnvName = os.Getenv("BRANCH_NAME")
	port = os.Getenv("BE_PORT")
	nftBackendUrl = os.Getenv("NFT_BACKEND_BASEURL")
}

func GetPort() string {
	LoadEnv()
	if port != "" {
		return ":" + port
	}
	return ":8000"
}

func GetNftBackendBaseUrl() string {
	LoadEnv()
	return nftBackendUrl
}
