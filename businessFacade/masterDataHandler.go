package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateMasterData(mdata model.MasterData) (string, error) {
	return masterdataRepository.CreateMasterData(mdata)
}

func CreateMasterDataCollection(data model.DataCollection) (string, error) {
	return masterdataRepository.CreateDataCollection(data)
}

func GetMasterDataByID(mDataID string) (model.MasterData, error) {
	return masterdataRepository.GetMasterDataByID(mDataID)
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
