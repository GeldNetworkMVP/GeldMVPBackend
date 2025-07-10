package services

import (
	"os"
	"strings"

	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
)

// Read the file from a specific location
func ReadFromFile(location string) string {
	content, err := os.ReadFile(location)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	readLine := strings.TrimSuffix(string(content), "\r\n")
	return readLine
}
