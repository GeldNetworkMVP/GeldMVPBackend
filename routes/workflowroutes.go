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
	//Will return Workflows based on Workflows ID provided
	model.Router{
		Name:    "Get Workflows by WorkflowID",
		Method:  "GET",
		Path:    "/workflows/{_id}",
		Handler: apiModel.GetWorkflowsByID,
	},
	//Will update Workflows based on Workflows ID provided
	model.Router{
		Name:    "Update Workflows",
		Method:  "PUT",
		Path:    "/updateworkflow",
		Handler: apiModel.UpdateWorkflow,
	},
}
