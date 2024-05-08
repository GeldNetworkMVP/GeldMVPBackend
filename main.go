package main

import (
	"net/http"
	"os"

	"github.com/GeldNetworkMVP/GeldMVPBackend/api/routes"
	"github.com/GeldNetworkMVP/GeldMVPBackend/commons"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
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
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Token"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := routes.NewRouter()
	// serve swagger documentation
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	// initial log file when server starts
	utilities.CreateLogFile()
	// create logger
	logger := utilities.NewCustomLogger()
	logger.LogWriter("Gateway Started @port "+port+" with "+envName+" environment", 1)

	http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
