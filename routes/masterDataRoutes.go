package routes

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/apiModel"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

var mDataRoutes = model.Routers{
	//Route will be used to add a new MasterData to DB
	model.Router{
		Name:    "Create MasterData",
		Method:  "POST",
		Path:    "/masterdata/save",
		Handler: apiModel.CreateMasterData,
	},
	//Will return MasterData based on MasterData ID provided
	model.Router{
		Name:    "Get MasterData by MasterDataID",
		Method:  "GET",
		Path:    "/masterdata/{_id}",
		Handler: apiModel.GetMasterDataByID,
	},
	//Will update MasterData based on MasterData ID provided
	model.Router{
		Name:    "Update MasterData",
		Method:  "PUT",
		Path:    "/updatemasterdata",
		Handler: apiModel.UpdateMasterData,
	},
}
