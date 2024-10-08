package dataConfigRepository

import (
	"context"
	"errors"

	"github.com/GeldNetworkMVP/GeldMVPBackend/database/connections"
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories"
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MasterDataRepository struct{}

var MasterData = "masterdata"
var DataCollection = "masterdata_records"

func (r *MasterDataRepository) CreateMasterData(mdata model.MasterData) (string, error) {
	return repositories.Save(mdata, MasterData)
}

func (r *MasterDataRepository) CreateDataCollection(data map[string]interface{}) (string, error) {
	return repositories.SaveDynamicData(data, DataCollection, "collectionname")
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

func (r *MasterDataRepository) GetRecordByID(ID string) (map[string]interface{}, error) {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetRecordByID : dataCollectionRepository: ", err.Error())
	}
	ctx := context.TODO()
	result := bson.M{}
	err = connections.GetSessionClient(DataCollection).FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Data Collection not found")
		}
		logs.ErrorLogger.Println("Error retrieving Data Collection:", err.Error())
		return nil, err
	}
	return result, nil
}

func (r *MasterDataRepository) GetRecordByMasterID(idName string, id string) ([]map[string]interface{}, error) {
	ctx := context.TODO()
	cursor, err := connections.GetSessionClient(DataCollection).Find(ctx, bson.M{idName: id})
	if err != nil {
		return nil, err
	}

	var collections []map[string]interface{}

	for cursor.Next(ctx) {
		var result map[string]interface{}
		err := cursor.Decode(&result)
		if err != nil {
			logs.ErrorLogger.Println("Error retrieving collections:", err.Error())
			return nil, err
		}
		collections = append(collections, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
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

func (r *MasterDataRepository) GetDataPaginatedResponse(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, data []map[string]interface{}, sort int) (model.DataPaginatedresponse, error) {
	contentResponse, paginationResponse, err := repositories.PaginateResponse[[]map[string]interface{}](
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

func (r *MasterDataRepository) UpdateDataCollection(UpdateObject requestDtos.UpdateDataCollection, update primitive.M) (map[string]interface{}, error) {
	var mDataUpdateResponse map[string]interface{}
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

func (r *MasterDataRepository) TestGetAllMasterData() ([]model.MasterData, error) {
	var allMasterData []model.MasterData
	findOptions := options.Find()
	result, err := connections.GetSessionClient(MasterData).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllMasterData:masterDataRepository.go: ", err.Error())
		return allMasterData, err
	}
	for result.Next(context.TODO()) {
		var masterdata model.MasterData
		err = result.Decode(&masterdata)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection partner in GetAllMasterData:masterDataRepository.go: ", err.Error())
			return allMasterData, err
		}
		allMasterData = append(allMasterData, masterdata)
	}
	return allMasterData, nil
}
