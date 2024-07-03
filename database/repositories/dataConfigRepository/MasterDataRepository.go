package dataConfigRepository

import (
	"context"

	"github.com/GeldNetworkMVP/GeldMVPBackend/database/connections"
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories"
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MasterDataRepository struct{}

var MasterData = "masterdata"
var DataCollection = "masterdata_records"

func (r *MasterDataRepository) CreateMasterData(mdata model.MasterData) (string, error) {
	return repositories.Save(mdata, MasterData)
}

func (r *MasterDataRepository) CreateDataCollection(data model.DataCollection) (string, error) {
	return repositories.Save(data, DataCollection)
}

func (r *MasterDataRepository) GetMasterDataByID(mDataID string) (model.MasterData, error) {
	var mdata model.MasterData
	objectId, err := primitive.ObjectIDFromHex(mDataID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetMasterDataByID : MasterDataRepository: ", err.Error())
	}
	rst, err := connections.GetSessionClient("masterdata").Find(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return mdata, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&mdata)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection document in GetMasterDataByID:MasterDataRepository.go: ", err.Error())
			return mdata, err
		}
	}
	return mdata, err
}

func (r *MasterDataRepository) GetRecordByID(ID string) (model.DataCollection, error) {
	var data model.DataCollection
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetRecordByID : MasterDataRepository: ", err.Error())
	}
	rst, err := connections.GetSessionClient("masterdata_records").Find(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return data, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&data)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection document in GetRecordByID:MasterDataRepository.go: ", err.Error())
			return data, err
		}
	}
	return data, err
}

func (r *MasterDataRepository) GetRecordByMasterID(idName string, id string) ([]model.DataCollection, error) {
	var collections []model.DataCollection
	rst, err := repositories.FindById(idName, id, DataCollection)
	if err != nil {
		return collections, err
	}
	for rst.Next(context.TODO()) {
		var collection model.DataCollection
		err = rst.Decode(&collection)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return collections, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (r *MasterDataRepository) GetMasterDataPaginatedResponse(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, mdata []model.MasterData, sort int) (model.MDataPaginatedresponse, error) {
	contentResponse, paginationResponse, err := repositories.PaginateResponse[[]model.MasterData](
		filterConfig,
		projectionData,
		pagesize,
		pageNo,
		collectionName,
		sortingFeildName,
		mdata,
		sort,
	)
	var response model.MDataPaginatedresponse
	if err != nil {
		return response, err
	}
	response.Content = contentResponse
	response.PaginationInfo = paginationResponse
	return response, nil
}

func (r *MasterDataRepository) GetDataPaginatedResponse(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, data []model.DataCollection, sort int) (model.DataPaginatedresponse, error) {
	contentResponse, paginationResponse, err := repositories.PaginateResponse[[]model.DataCollection](
		filterConfig,
		projectionData,
		pagesize,
		pageNo,
		collectionName,
		sortingFeildName,
		data,
		sort,
	)
	var response model.DataPaginatedresponse
	if err != nil {
		return response, err
	}
	response.Content = contentResponse
	response.PaginationInfo = paginationResponse
	return response, nil
}

func (r *MasterDataRepository) UpdateMasterData(UpdateObject requestDtos.UpdateMasterData, update primitive.M) (model.MasterData, error) {
	var mDataUpdateResponse model.MasterData
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("masterdata").FindOneAndUpdate(context.TODO(), bson.M{"_id": UpdateObject.DataID}, update, &opt)
	if rst != nil {
		err := rst.Decode((&mDataUpdateResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error Occured while Update MasterData", err.Error())
			return mDataUpdateResponse, err
		}
		return mDataUpdateResponse, err
	}
	return mDataUpdateResponse, nil
}

func (r *MasterDataRepository) UpdateDataCollection(UpdateObject requestDtos.UpdateDataCollection, update primitive.M) (model.DataCollection, error) {
	var mDataUpdateResponse model.DataCollection
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("masterdata_records").FindOneAndUpdate(context.TODO(), bson.M{"_id": UpdateObject.CollectionID}, update, &opt)
	if rst != nil {
		err := rst.Decode((&mDataUpdateResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error Occured while Update MasterDataCollection", err.Error())
			return mDataUpdateResponse, err
		}
		return mDataUpdateResponse, err
	}
	return mDataUpdateResponse, nil
}

func (r *MasterDataRepository) DeleteMasterData(DataID primitive.ObjectID) error {
	result, err := connections.GetSessionClient(MasterData).DeleteOne(context.TODO(), bson.M{"_id": DataID})
	if err != nil {
		logs.ErrorLogger.Println("Error occured when Connecting to DB and executing DeleteOne Query in DeleteMasterData(MasterDataRepository): ", err.Error())
	}
	logs.InfoLogger.Println("MasterData deleted :", result.DeletedCount)
	return err

}

func (r *MasterDataRepository) DeleteMasterDataRecords(RecordID primitive.ObjectID) error {
	result, err := connections.GetSessionClient(DataCollection).DeleteOne(context.TODO(), bson.M{"_id": RecordID})
	if err != nil {
		logs.ErrorLogger.Println("Error occured when Connecting to DB and executing DeleteOne Query in DeleteDataCollection(MasterDataRepository): ", err.Error())
	}
	logs.InfoLogger.Println("DataCollection deleted :", result.DeletedCount)
	return err

}
