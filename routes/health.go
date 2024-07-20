package routes

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/apiModel"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

// This routes use to check the API status

var HealthRoutes = model.Routers{

	model.Router{
		Name:    "Connection test API",
		Method:  "GET",
		Path:    "/api/health",
		Handler: apiModel.HealthCheck,
	},
}
