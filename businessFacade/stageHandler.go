package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
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

func DeleteStageByID(StageID string) error {
	return stageRepository.DeleteStage(StageID)
}

func GetStageDataPagination(paginationData requestDtos.StagesForMatrixView) (model.StagePaginatedresponse, error) {
	filter := bson.M{
		"userid": paginationData.UserID,
	}
	projection := GetProjectionDataMatrixViewForStageData()
	var data []model.Stages
	response, err := stageRepository.GetStageDataPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "stages", "userid", data, paginationData.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return model.StagePaginatedresponse(response), err
	}
	return model.StagePaginatedresponse(response), err
}

func GetProjectionDataMatrixViewForStageData() bson.D {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "userid", Value: 1},
		{Key: "stagename", Value: 1},
		{Key: "description", Value: 1},
		{Key: "fields", Value: 1},
	}
	return projection
}
