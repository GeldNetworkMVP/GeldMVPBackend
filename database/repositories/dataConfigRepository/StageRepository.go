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

type StageRepository struct{}

var Stages = "stages"

func (r *StageRepository) CreateStages(stages model.Stages) (string, error) {
	return repositories.Save(stages, Stages)
}

func (r *StageRepository) GetStagesByID(stageID string) (model.Stages, error) {
	var stages model.Stages
	objectId, err := primitive.ObjectIDFromHex(stageID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetStagesByID : stageRepository: ", err.Error())
	}
	rst, err := connections.GetSessionClient("stages").Find(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return stages, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&stages)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection document in GetStageByID:stageRepository.go: ", err.Error())
			return stages, err
		}
	}
	return stages, err
}

func (r *StageRepository) UpdateStages(UpdateObject requestDtos.UpdateStages, update primitive.M) (model.Stages, error) {
	var StagesUpdateResponse model.Stages
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("stages").FindOneAndUpdate(context.TODO(), bson.M{"_id": UpdateObject.StageID}, update, &opt)
	if rst != nil {
		err := rst.Decode((&StagesUpdateResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error Occured while Update Stages", err.Error())
			return StagesUpdateResponse, err
		}
		return StagesUpdateResponse, err
	}
	return StagesUpdateResponse, nil
}

func (r *StageRepository) DeleteStage(stageID primitive.ObjectID) error {
	result, err := connections.GetSessionClient(Stages).DeleteOne(context.TODO(), bson.M{"_id": stageID})
	if err != nil {
		logs.ErrorLogger.Println("Error occured when Connecting to DB and executing DeleteOne Query in DeleteStage(stageRepository): ", err.Error())
	}
	logs.InfoLogger.Println("stage deleted :", result.DeletedCount)
	return err

}

func (r *StageRepository) GetStageDataPaginatedResponse(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, data []model.Stages, sort int) (model.StagePaginatedresponse, error) {
	contentResponse, paginationResponse, err := repositories.PaginateResponse[[]model.Stages](
		filterConfig,
		projectionData,
		pagesize,
		pageNo,
		collectionName,
		sortingFeildName,
		data,
		sort,
	)
	var response model.StagePaginatedresponse
	if err != nil {
		return response, err
	}
	response.Content = contentResponse
	response.PaginationInfo = paginationResponse
	return response, nil
}

func (r *StageRepository) GetStagesByName(stageName string) (model.Stages, error) {
	var stages model.Stages

	rst, err := connections.GetSessionClient("stages").Find(context.TODO(), bson.M{"stagename": stageName})
	if err != nil {
		return stages, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&stages)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection document in GetStageByName:stageRepository.go: ", err.Error())
			return stages, err
		}
	}
	return stages, err
}

func (r *StageRepository) TestGetAllStages() ([]model.Stages, error) {
	var allStages []model.Stages
	findOptions := options.Find()
	result, err := connections.GetSessionClient(Stages).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllStages:StagesRepository.go: ", err.Error())
		return allStages, err
	}
	for result.Next(context.TODO()) {
		var stage model.Stages
		err = result.Decode(&stage)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection partner in GetAllStages:StagesRepository.go: ", err.Error())
			return allStages, err
		}
		allStages = append(allStages, stage)
	}
	return allStages, nil
}
