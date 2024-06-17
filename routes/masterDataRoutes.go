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
	//Route will be used to add a new DataCollection to DB
	model.Router{
		Name:    "Create DataCollection",
		Method:  "POST",
		Path:    "/record/save",
		Handler: apiModel.CreateDataCollection,
	},
	//Will return MasterData based on MasterData ID provided
	model.Router{
		Name:    "Get MasterData by MasterDataID",
		Method:  "GET",
		Path:    "/masterdata/{_id}",
		Handler: apiModel.GetMasterDataByID,
	},
	//Will return Records based on  ID provided
	model.Router{
		Name:    "Get RecordData by ID",
		Method:  "GET",
		Path:    "/record/{_id}",
		Handler: apiModel.GetRecordDataByID,
	},
	//Will return Records based on MasterData ID provided
	model.Router{
		Name:    "Get RecordData by mDataID",
		Method:  "GET",
		Path:    "/record/{dataid}",
		Handler: apiModel.GetRecordDataByMasterDataID,
	},
	//Will return user based masterdata paginated response
	model.Router{
		Name:    "Get Master data pagination",
		Method:  "Get",
		Path:    "/masterdata/{userid}",
		Handler: apiModel.GetPaginatedMasterData,
	},
	//Will return master data based record paginated response
	model.Router{
		Name:    "Get Master data record pagination",
		Method:  "Get",
		Path:    "/record/{dataid}",
		Handler: apiModel.GetPaginatedData,
	},
	//Will update MasterData based on MasterData ID provided
	model.Router{
		Name:    "Update MasterData",
		Method:  "PUT",
		Path:    "/updatemasterdata",
		Handler: apiModel.UpdateMasterData,
	},
	//Will delete MasterData based on MasterData ID provided
	model.Router{
		Name:    "Delete MasterData by MasterDataID",
		Method:  "DELETE",
		Path:    "/masterdata/remove/{_id}",
		Handler: apiModel.DeleteMasterDataByID,
	},
	//Will delete Record based on Data ID provided
	model.Router{
		Name:    "Delete Record by DataID",
		Method:  "DELETE",
		Path:    "/record/remove/{_id}",
		Handler: apiModel.DeleteMasterDataRecordByID,
	},
}
