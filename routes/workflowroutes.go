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
	//Will delete Workflow based on Workflow ID provided
	model.Router{
		Name:    "Delete Workflow by WorkflowID",
		Method:  "DELETE",
		Path:    "/workflows/remove/{_id}",
		Handler: apiModel.DeleteWorkflowByID,
	},
	//Will return workflows based userid paginated response
	// model.Router{
	// 	Name:    "Get workflow data pagination",
	// 	Method:  "Get",
	// 	Path:    "/userworkflows/{userid}",
	// 	Handler: apiModel.GetPaginatedWorkflowData,
	// },

	//Will return workflows based userid paginated response
	model.Router{
		Name:    "Test workflow data pagination",
		Method:  "Get",
		Path:    "/paginatedworkflows",
		Handler: apiModel.TestPaginatedWorkflowData,
	},
	//Will return all the Workflows
	model.Router{
		Name:    "Test normal Get All Workflows",
		Method:  "GET",
		Path:    "/workflows/",
		Handler: apiModel.TestGetAllWorkflows,
	},
}
