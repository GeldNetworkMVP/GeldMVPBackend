package userAndPermissions

import (
	"context"
	"fmt"

	"github.com/GeldNetworkMVP/GeldMVPBackend/database/connections"
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories"
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct{}

var User = "appusers"

func (r *UserRepository) CreateUsers(users model.AppUser) (string, error) {
	return repositories.Save(users, User)
}

func (r *UserRepository) GetUserByID(userID string) (model.AppUser, error) {
	var user model.AppUser
	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetUserByID : userRepository: ", err.Error())
	}
	rst, err := connections.GetSessionClient("appusers").Find(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return user, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&user)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection document in GetUserByID:userRepository.go: ", err.Error())
			return user, err
		}
	}
	return user, err
}

func (r *UserRepository) UpdateUsers(UpdateObject requestDtos.UpdateUser, update primitive.M) (model.AppUser, error) {
	var appUpdateResponse model.AppUser
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("appusers").FindOneAndUpdate(context.TODO(), bson.M{"_id": UpdateObject.AppUserID}, update, &opt)
	if rst != nil {
		err := rst.Decode((&appUpdateResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error Occured while Update User", err.Error())
			return appUpdateResponse, err
		}
		return appUpdateResponse, err
	}
	return appUpdateResponse, nil
}

func (r *UserRepository) DeleteUser(userID primitive.ObjectID) error {
	fmt.Println("id", userID)
	result, err := connections.GetSessionClient(User).DeleteOne(context.TODO(), bson.M{"_id": userID})
	if err != nil {
		logs.ErrorLogger.Println("Error occured when Connecting to DB and executing DeleteOne Query in DeleteUser(userRepository): ", err.Error())
	}
	logs.InfoLogger.Println("user deleted :", result.DeletedCount)
	return err

}

func (r *UserRepository) GetUserssDataPaginatedResponse(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, data []model.AppUser, sort int) (model.UserPaginatedResponse, error) {
	contentResponse, paginationResponse, err := repositories.PaginateResponse[[]model.AppUser](
		filterConfig,
		projectionData,
		pagesize,
		pageNo,
		collectionName,
		sortingFeildName,
		data,
		sort,
	)
	var response model.UserPaginatedResponse
	if err != nil {
		return response, err
	}
	response.Content = contentResponse
	response.PaginationInfo = paginationResponse
	return response, nil
}
