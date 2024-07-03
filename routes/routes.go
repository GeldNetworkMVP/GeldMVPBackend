package routes

import "github.com/GeldNetworkMVP/GeldMVPBackend/model"

var ApplicationRoutes model.Routers

func init() {
	routes := []model.Routers{
		WorkflowRoutes,
		StageRoutes,
		HealthRoutes,
		MDataRoutes,
		DataTemplateRoutes,
		AppUserRoutes,
	}

	for _, r := range routes {
		ApplicationRoutes = append(ApplicationRoutes, r...)
	}
}
