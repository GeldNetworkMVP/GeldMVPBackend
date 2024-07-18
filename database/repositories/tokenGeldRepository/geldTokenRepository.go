package tokenGeldRepository

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

type TokenRepository struct{}

var Tokens = "tokens"
var Transactions = "transactions"

func (r *TokenRepository) SaveTokens(tokens model.Tokens) (string, error) {
	return repositories.Save(tokens, Tokens)
}

func (r *TokenRepository) SaveTransactions(transactions model.TokenTransactions) (string, error) {
	return repositories.Save(transactions, Transactions)
}

func (r *TokenRepository) GetTokensByID(tokenID string) (model.Tokens, error) {
	var tokens model.Tokens
	objectId, err := primitive.ObjectIDFromHex(tokenID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetTokensByID : tokenRepository: ", err.Error())
	}
	rst, err := connections.GetSessionClient(Tokens).Find(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return tokens, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&tokens)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection document in GetTokenByID:tokenRepository.go: ", err.Error())
			return tokens, err
		}
	}
	return tokens, err
}

func (r *TokenRepository) UpdateTokens(UpdateObject requestDtos.UpdateToken, update primitive.M) (model.Tokens, error) {
	var TokenUpdateResponse model.Tokens
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient(Tokens).FindOneAndUpdate(context.TODO(), bson.M{"_id": UpdateObject.TokenID}, update, &opt)
	if rst != nil {
		err := rst.Decode((&TokenUpdateResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error Occured while Update Token", err.Error())
			return TokenUpdateResponse, err
		}
		return TokenUpdateResponse, err
	}
	return TokenUpdateResponse, nil
}

func (r *TokenRepository) GetTokensPaginatedResponseByStatus(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, data []model.Tokens, sort int) (model.TokenPaginatedresponse, error) {
	contentResponse, paginationResponse, err := repositories.PaginateResponse[[]model.Tokens](
		filterConfig,
		projectionData,
		pagesize,
		pageNo,
		collectionName,
		sortingFeildName,
		data,
		sort,
	)
	var response model.TokenPaginatedresponse
	if err != nil {
		return response, err
	}
	response.Content = contentResponse
	response.PaginationInfo = paginationResponse
	return response, nil
}
