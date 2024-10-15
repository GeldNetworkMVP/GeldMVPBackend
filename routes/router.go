package routes

import (
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/commons"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(commons.JWTValidationMiddleware)
	for _, route := range ApplicationRoutes {

		var handler http.Handler = route.Handler

		if route.Protected {
			protectedRoutes.
				Methods(route.Method).
				Path(route.Path).
				Name(route.Name).
				Handler(handler)
		} else {
			// Otherwise, add it directly to the main router.
			router.
				Methods(route.Method).
				Path(route.Path).
				Name(route.Name).
				Handler(handler)
		}

	}

	return router
}
