package routes

import (
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/api/businessFacade"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes An Array of type Route
type Routes []Route

/*
routes contains all the routes
@author Azeem Ashraf, Jajeththanan Sabapathipillai
*/
var routes = Routes{
	Route{
		"Get server health",
		"GET",
		"/health",
		businessFacade.HealthCheck,
	},
}
