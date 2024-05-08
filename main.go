package main

import (
	"os"

	"github.com/GeldNetworkMVP/GeldMVPBackend/commons"
)

func getPort() string {
	p := os.Getenv("GATEWAY_PORT")
	if p != "" {
		return ":" + p
	}
	return ":8000"
}

func main() {
	// godotenv package
	envName := commons.GoDotEnvVariable("ENVIRONMENT")

	// getEnvironment()
	port := getPort()
}
