package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateStages(stages model.Stages) (string, error) {
	return stageRepository.CreateStages(stages)
}

func GetStagesByID(stageID string) (model.Stages, error) {
	return stageRepository.GetStagesByID(stageID)
}

func UpdateStages(UpdateObject requestDtos.UpdateStages) (model.Stages, error) {
	update := bson.M{
		"$set": bson.M{"stagename": UpdateObject.StageName, "description": UpdateObject.Description, "fields": UpdateObject.Fields},
	}
	return stageRepository.UpdateStages(UpdateObject, update)
}
