package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUsers(users model.AppUser) (string, error) {
	return userRepository.CreateUsers(users)
}

func GetUserByID(userID string) (model.AppUser, error) {
	return userRepository.GetUserByID(userID)
}

func UpdateUsers(UpdateObject requestDtos.UpdateUser) (model.AppUser, error) {
	update := bson.M{
		"$set": bson.M{"username": UpdateObject.Username, "email": UpdateObject.Email, "contact": UpdateObject.Contact, "designation": UpdateObject.Designation, "encpw": UpdateObject.EncPW, "status": UpdateObject.Status},
	}
	return userRepository.UpdateUsers(UpdateObject, update)
}

func DeleteUserByID(userID primitive.ObjectID) error {
	return userRepository.DeleteUser(userID)
}

func GetUserDataPagination(paginationData requestDtos.UserForMatrixView) (model.UserPaginatedResponse, error) {
	filter := bson.M{
		"userid": paginationData.UserID,
	}
	projection := GetProjectionDataMatrixViewForUserData()
	var data []model.AppUser
	response, err := userRepository.GetUserssDataPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "appusers", "userid", data, paginationData.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return model.UserPaginatedResponse(response), err
	}
	return model.UserPaginatedResponse(response), err
}

func GetProjectionDataMatrixViewForUserData() bson.D {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "username", Value: 1},
		{Key: "email", Value: 1},
		{Key: "contact", Value: 1},
		{Key: "designation", Value: 1},
		{Key: "status", Value: 1},
	}
	return projection
}
