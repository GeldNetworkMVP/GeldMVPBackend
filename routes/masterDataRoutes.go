package routes

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/apiModel"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

var MDataRoutes = model.Routers{
	//Route will be used to add a new MasterData to DB
	model.Router{
		Name:      "Create MasterData",
		Method:    "POST",
		Path:      "/masterdata/save",
		Handler:   apiModel.CreateMasterData,
		Protected: true,
	},
	//Route will be used to add a new DataCollection to DB
	model.Router{
		Name:      "Create DataCollection",
		Method:    "POST",
		Path:      "/record/save",
		Handler:   apiModel.CreateDataCollection,
		Protected: true,
	},
	//Will return MasterData based on MasterData ID provided
	model.Router{
		Name:      "Get MasterData by MasterDataID",
		Method:    "GET",
		Path:      "/masterdata/{_id}",
		Handler:   apiModel.GetMasterDataByID,
		Protected: true,
	},
	//Will return Records based on  ID provided
	model.Router{
		Name:      "Get RecordData by ID",
		Method:    "GET",
		Path:      "/record/{_id}",
		Handler:   apiModel.GetRecordDataByID,
		Protected: true,
	},
	//Will return Records based on MasterData ID provided
	model.Router{
		Name:      "Get RecordData by mDataID",
		Method:    "GET",
		Path:      "/records/{dataid}",
		Handler:   apiModel.GetRecordDataByMasterDataID,
		Protected: true,
	},
	//Will return user based masterdata paginated response
	// model.Router{
	// 	Name:    "Get Master data pagination",
	// 	Method:  "Get",
	// 	Path:    "/usermasterdata/{userid}",
	// 	Handler: apiModel.GetPaginatedMasterData,
	// },
	//Will return record paginated response
	model.Router{
		Name:      "Get  record pagination",
		Method:    "Get",
		Path:      "/masterrecord/{dataid}",
		Handler:   apiModel.GetPaginatedData,
		Protected: true,
	},
	//Will update MasterData based on MasterData ID provided
	model.Router{
		Name:      "Update MasterData",
		Method:    "PUT",
		Path:      "/updatemasterdata",
		Handler:   apiModel.UpdateMasterData,
		Protected: true,
	},
	//Will update MasterDataRecords based on Collection ID provided
	model.Router{
		Name:      "Update Records",
		Method:    "PUT",
		Path:      "/updaterecords",
		Handler:   apiModel.UpdateDataCollection,
		Protected: true,
	},
	//Will delete MasterData based on MasterData ID provided
	model.Router{
		Name:      "Delete MasterData by MasterDataID",
		Method:    "DELETE",
		Path:      "/masterdata/remove/{_id}",
		Handler:   apiModel.DeleteMasterDataByID,
		Protected: true,
	},
	//Will delete Record based on Data ID provided
	model.Router{
		Name:      "Delete Record by DataID",
		Method:    "DELETE",
		Path:      "/record/remove/{_id}",
		Handler:   apiModel.DeleteMasterDataRecordByID,
		Protected: true,
	},
	//Get all Plots for Plot ID
	model.Router{
		Name:      "Get plot data record by container ID",
		Method:    "Get",
		Path:      "/plotrecord",
		Handler:   apiModel.GetPlotDataByMasterDataID,
		Protected: true,
	},
	//Will return all the Master Data Containers
	model.Router{
		Name:      "Test normal Get All Master Data",
		Method:    "GET",
		Path:      "/masterdata",
		Handler:   apiModel.TestGetAllMasterData,
		Protected: true,
	},
}
