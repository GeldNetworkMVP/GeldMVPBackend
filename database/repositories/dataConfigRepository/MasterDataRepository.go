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
