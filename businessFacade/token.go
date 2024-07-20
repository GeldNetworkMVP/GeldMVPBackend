package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveTokens(tokens model.Tokens) (string, error) {
	return tokensRepository.SaveTokens(tokens)
}

func SaveTransaction(transaction model.TokenTransactions) (string, error) {
	return tokensRepository.SaveTransactions(transaction)
}

func GetTokenByID(tokensID string) (model.Tokens, error) {
	return tokensRepository.GetTokensByID(tokensID)
}

func UpdateToken(UpdateObject requestDtos.UpdateToken) (model.Tokens, error) {
	update := bson.M{
		"$set": bson.M{"status": UpdateObject.Status},
	}
	return tokensRepository.UpdateTokens(UpdateObject, update)
}

func GetTokenPaginationByStatus(paginationData requestDtos.TokenForMatrixView) (model.TokenPaginatedresponse, error) {
	filter := bson.M{
		"status": paginationData.Status,
	}
	projection := GetProjectionDataMatrixViewForTokenData()
	var data []model.Tokens
	response, err := tokensRepository.GetTokensPaginatedResponseByStatus(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "tokens", "status", data, paginationData.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return model.TokenPaginatedresponse(response), err
	}
	return model.TokenPaginatedresponse(response), err
}

func GetProjectionDataMatrixViewForTokenData() bson.D {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "plotid", Value: 1},
		{Key: "tokenname", Value: 1},
		{Key: "description", Value: 1},
		{Key: "cid", Value: 1},
		{Key: "price", Value: 1},
		{Key: "status", Value: 1},
		{Key: "bcstatus", Value: 1},
		{Key: "tokenhash", Value: 1},
	}
	return projection
}
