package routes

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/apiModel"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

var WorkflowRoutes = model.Routers{
	//Route will be used to add a new Workflow to DB
	model.Router{
		Name:    "Create Workflows",
		Method:  "POST",
		Path:    "/workflows/save",
		Handler: apiModel.CreateWorkflow,
	},
}
