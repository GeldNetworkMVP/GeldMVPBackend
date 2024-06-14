package routes

import "github.com/GeldNetworkMVP/GeldMVPBackend/model"

var ApplicationRoutes model.Routers

func init() {
	routes := []model.Routers{
		WorkflowRoutes,
	}

	for _, r := range routes {
		ApplicationRoutes = append(ApplicationRoutes, r...)
	}
}
