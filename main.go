package main

import (
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/configs"
	"github.com/GeldNetworkMVP/GeldMVPBackend/routes"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
	//"github.com/joho/godotenv"
	//"github.com/sirupsen/logrus"
)

// @title			Geld.Network API
// @version		1.0
// @description	This is the Geld.Network Server.
// @termsOfService	http://swagger.io/terms/
func main() {
	logs.InfoLogger.Println("Tracified Backend")
	// err := godotenv.Load()
	// if err != nil {
	// 	logrus.Println("Info Issue with loading .env file")
	// 	logs.InfoLogger.Println("Info Issue with loading .env1 file")
	// }
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Token"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	// Start API
	router := routes.NewRouter()

	router.Handle("/docs/swagger.yaml", http.FileServer(http.Dir("./")))
	opts := middleware.SwaggerUIOpts{SpecURL: "/docs/swagger.yaml", Path: "api-docs"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/api-docs", sh)

	http.Handle("/api/", router)
	logs.InfoLogger.Println("Gateway Started @port " + configs.GetPort() + " with " + configs.EnvName + " environment")
	http.ListenAndServe(configs.GetPort(), handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
