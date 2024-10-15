package routes

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/apiModel"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

var StageRoutes = model.Routers{
	//Route will be used to add a new Stage to DB
	model.Router{
		Name:      "Create Stage",
		Method:    "POST",
		Path:      "/stage/save",
		Handler:   apiModel.CreateStage,
		Protected: true,
	},
	//Will return Stages based on Stages ID provided
	model.Router{
		Name:      "Get Stages by StageID",
		Method:    "GET",
		Path:      "/stage/{_id}",
		Handler:   apiModel.GetStagesByID,
		Protected: true,
	},
	//Will update Stages based on Stage ID provided
	model.Router{
		Name:      "Update Stages",
		Method:    "PUT",
		Path:      "/updatestage",
		Handler:   apiModel.UpdateStages,
		Protected: true,
	},
	//Will delete Stage based on Stage ID provided
	model.Router{
		Name:      "Delete Stage by StageID",
		Method:    "DELETE",
		Path:      "/stage/remove/{_id}",
		Handler:   apiModel.DeleteStageByID,
		Protected: true,
	},
	//Will return stages based userid paginated response
	// model.Router{
	// 	Name:    "Get stage data  pagination",
	// 	Method:  "Get",
	// 	Path:    "/userstage/{userid}",
	// 	Handler: apiModel.GetPaginatedStageData,
	// },
	//Will return Stages based on Stages Name provided
	model.Router{
		Name:      "Get Stages by Stage Name",
		Method:    "GET",
		Path:      "/stagename/{stagename}",
		Handler:   apiModel.GetStagesByName,
		Protected: true,
	},
	//Will return all the Stages
	model.Router{
		Name:      "Test normal Get All Stages",
		Method:    "GET",
		Path:      "/stages",
		Handler:   apiModel.TestGetAllStages,
		Protected: true,
	},
}
