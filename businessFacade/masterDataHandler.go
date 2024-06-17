package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateMasterData(mdata model.MasterData) (string, error) {
	return masterdataRepository.CreateMasterData(mdata)
}

func CreateMasterDataCollection(data model.DataCollection) (string, error) {
	return masterdataRepository.CreateDataCollection(data)
}

func GetRecordDataByID(ID string) (model.DataCollection, error) {
	return masterdataRepository.GetRecordByID(ID)
}

func GetRecordDataByMasterDataID(mDataID string) ([]model.DataCollection, error) {
	return masterdataRepository.GetRecordByMasterID("dataid", mDataID)
}
func GetMasterDataByID(mDataID string) (model.MasterData, error) {
	return masterdataRepository.GetMasterDataByID(mDataID)
}

func GetMasterDataPagination(paginationData requestDtos.MasterDataForMatrixView) (model.MDataPaginatedresponse, error) {
	filter := bson.M{
		"userid": paginationData.UserID,
	}
	projection := GetProjectionDataMatrixViewForMasterData()
	var mdata []model.MasterData
	response, err := masterdataRepository.GetMasterDataPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "masterdata", "userid", mdata, paginationData.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return model.MDataPaginatedresponse(response), err
	}
	return model.MDataPaginatedresponse(response), err
}

func GetProjectionDataMatrixViewForMasterData() bson.D {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "userid", Value: 1},
		{Key: "dataname", Value: 1},
		{Key: "description", Value: 1},
		{Key: "dataCollection", Value: 1},
	}
	return projection
}

func GetDataPagination(paginationData requestDtos.DataRecordForMatrixView) (model.DataPaginatedresponse, error) {
	filter := bson.M{
		"dataid": paginationData.DataID,
	}
	projection := GetProjectionDataMatrixViewForData()
	var data []model.DataCollection
	response, err := masterdataRepository.GetDataPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "masterdata_records", "dataid", data, paginationData.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return model.DataPaginatedresponse(response), err
	}
	return model.DataPaginatedresponse(response), err
}

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
		"$set": bson.M{"dataname": UpdateObject.DataName, "description": UpdateObject.Description},
	}
	return masterdataRepository.UpdateMasterData(UpdateObject, update)
}

func UpdateDataCollection(UpdateObject requestDtos.UpdateDataCollection) (model.DataCollection, error) {
	update := bson.M{
		"$set": bson.M{"DataID": UpdateObject.DataID,
			"collectionname": UpdateObject.CollectionName,
			"description":    UpdateObject.Description,
			"purpose":        UpdateObject.Purpose,
			"location":       UpdateObject.Location,
			"contact":        UpdateObject.Contact,
			"type":           UpdateObject.Type},
	}
	return masterdataRepository.UpdateDataCollection(UpdateObject, update)
}

func DeleteMasterDataByID(mDataID string) error {
	return masterdataRepository.DeleteMasterData(mDataID)
}

func DeleteMasterDataRecordByID(DataID string) error {
	return masterdataRepository.DeleteMasterDataRecords(DataID)
}
