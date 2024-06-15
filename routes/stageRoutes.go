package routes

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/apiModel"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

var StageRoutes = model.Routers{
	//Route will be used to add a new Stage to DB
	model.Router{
		Name:    "Create Stage",
		Method:  "POST",
		Path:    "/stage/save",
		Handler: apiModel.CreateStage,
	},
	//Will return Stages based on Stages ID provided
	model.Router{
		Name:    "Get Stages by StageID",
		Method:  "GET",
		Path:    "/stage/{_id}",
		Handler: apiModel.GetStagesByID,
	},
	//Will update Stages based on Stage ID provided
	model.Router{
		Name:    "Update Stages",
		Method:  "PUT",
		Path:    "/updatestage",
		Handler: apiModel.UpdateStages,
	},
}
