package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMasterData(mdata model.MasterData) (string, error) {
	return masterdataRepository.CreateMasterData(mdata)
}

func CreateMasterDataCollection(data map[string]interface{}) (string, error) {
	return masterdataRepository.CreateDataCollection(data)
}

func GetRecordDataByID(ID string) (map[string]interface{}, error) {
	return masterdataRepository.GetRecordByID(ID)
}

func GetRecordDataByMasterDataID(mDataID string) ([]map[string]interface{}, error) {
	return masterdataRepository.GetRecordByMasterID("dataid", mDataID)
}
func GetMasterDataByID(mDataID string) (model.MasterData, error) {
	return masterdataRepository.GetMasterDataByID(mDataID)
}

// func GetMasterDataPagination(paginationData requestDtos.MasterDataForMatrixView) (model.MDataPaginatedresponse, error) {
// 	filter := bson.M{
// 		"userid": paginationData.UserID,
// 	}
// 	projection := GetProjectionDataMatrixViewForMasterData()
// 	var mdata []model.MasterData
// 	response, err := masterdataRepository.GetMasterDataPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "masterdata", "userid", mdata, paginationData.SortType)
// 	if err != nil {
// 		logs.ErrorLogger.Println("Error occurred :", err.Error())
// 		return model.MDataPaginatedresponse(response), err
// 	}
// 	return model.MDataPaginatedresponse(response), err
// }

func GetProjectionDataMatrixViewForMasterData() bson.D {
	projection := bson.D{
		{Key: "_id", Value: 1},
		// {Key: "userid", Value: 1},
		{Key: "dataname", Value: 1},
		{Key: "description", Value: 1},
		{Key: "fields", Value: 1},
	}
	return projection
}

// func GetDataPagination(paginationData requestDtos.DataRecordForMatrixView) (model.DataPaginatedresponse, error) {
// 	filter := bson.M{
// 		"dataid": paginationData.DataID,
// 	}
// 	projection := GetProjectionDataMatrixViewForData()
// 	var data []model.DataCollection
// 	response, err := masterdataRepository.GetDataPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "masterdata_records", "dataid", data, paginationData.SortType)
// 	if err != nil {
// 		logs.ErrorLogger.Println("Error occurred :", err.Error())
// 		return model.DataPaginatedresponse(response), err
// 	}
// 	return model.DataPaginatedresponse(response), err
// }

func GetProjectionDataMatrixViewForData() bson.D {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "dataid", Value: 1},
		{Key: "userid", Value: 1},
		{Key: "collectionname", Value: 1},
		{Key: "description", Value: 1},
		{Key: "purpose", Value: 1},
		{Key: "location", Value: 1},
		{Key: "contact", Value: 1},
		{Key: "type", Value: 1},
	}
	return projection
}

func UpdateMasterData(UpdateObject requestDtos.UpdateMasterData) (model.MasterData, error) {
	update := bson.M{
		"$set": bson.M{"dataname": UpdateObject.DataName, "description": UpdateObject.Description, "fields": UpdateObject.Fields},
	}
	return masterdataRepository.UpdateMasterData(UpdateObject, update)
}

// func UpdateDataCollection(UpdateObject requestDtos.UpdateDataCollection) (model.DataCollection, error) {
// 	update := bson.M{
// 		"$set": bson.M{"DataID": UpdateObject.DataID,
// 			"collectionname": UpdateObject.CollectionName,
// 			"description":    UpdateObject.Description,
// 			"purpose":        UpdateObject.Purpose,
// 			"location":       UpdateObject.Location,
// 			"contact":        UpdateObject.Contact,
// 			"type":           UpdateObject.Type},
// 	}
// 	return masterdataRepository.UpdateDataCollection(UpdateObject, update)
// }

func UpdateDataCollection(updateData requestDtos.UpdateDataCollection) (map[string]interface{}, error) {
	update := bson.M{
		"$set": updateData.DataObject,
	}

	return masterdataRepository.UpdateDataCollection(updateData, update)
}

func DeleteMasterDataByID(mDataID primitive.ObjectID) error {
	return masterdataRepository.DeleteMasterData(mDataID)
}

func DeleteMasterDataRecordByID(DataID primitive.ObjectID) error {
	return masterdataRepository.DeleteMasterDataRecords(DataID)
}

func GetDataPagination(paginationData requestDtos.DataRecordForMatrixView) (model.DataPaginatedresponse, error) {
	filter := bson.M{
		"dataid": paginationData.DataID,
	}
	projection := GetDynamicProjectionForData(paginationData.Fields)

	var data []map[string]interface{}

	response, err := masterdataRepository.GetDataPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "masterdata_records", "dataid", data, paginationData.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred:", err.Error())
		return model.DataPaginatedresponse{}, err
	}

	return model.DataPaginatedresponse(response), nil
}

func GetDynamicProjectionForData(fields []string) bson.D {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "dataid", Value: 1},
	}

	for _, field := range fields {
		projection = append(projection, bson.E{Key: field, Value: 1})
	}
	return projection
}

func TestGetAllMasterData() ([]model.MasterData, error) {
	return masterdataRepository.TestGetAllMasterData()
}
