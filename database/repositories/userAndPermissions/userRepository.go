package userAndPermissions

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

type UserRepository struct{}

var User = "appusers"

func (r *UserRepository) CreateUsers(users model.AppUser) (string, error) {
	return repositories.Save(users, User)
}

func (r *UserRepository) GetUserByID(userID string) (model.AppUserDetails, error) {
	var user model.AppUserDetails
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

func (r *UserRepository) UpdateUsers(UpdateObject requestDtos.UpdateUser, update primitive.M) (model.AppUserDetails, error) {
	var appUpdateResponse model.AppUserDetails
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
	result, err := connections.GetSessionClient(User).DeleteOne(context.TODO(), bson.M{"_id": userID})
	if err != nil {
		logs.ErrorLogger.Println("Error occured when Connecting to DB and executing DeleteOne Query in DeleteUser(userRepository): ", err.Error())
	}
	logs.InfoLogger.Println("user deleted :", result.DeletedCount)
	return err

}

func (r *UserRepository) GetUserssDataPaginatedResponse(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, data []model.AppUserDetails, sort int) (model.UserPaginatedResponse, error) {
	contentResponse, paginationResponse, err := repositories.PaginateResponse[[]model.AppUserDetails](
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

func (r *UserRepository) TestGetAllUsers() ([]model.AppUserDetails, error) {
	var allUsers []model.AppUserDetails
	findOptions := options.Find()
	result, err := connections.GetSessionClient(User).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllUsers:UsersRepository.go: ", err.Error())
		return allUsers, err
	}
	for result.Next(context.TODO()) {
		var user model.AppUserDetails
		err = result.Decode(&user)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection partner in GetAllUserss:UsersRepository.go: ", err.Error())
			return allUsers, err
		}
		allUsers = append(allUsers, user)
	}
	return allUsers, nil
}

func (r *UserRepository) GetUsersByField(idName string, id string) ([]model.AppUserDetails, error) {
	ctx := context.TODO()
	cursor, err := connections.GetSessionClient(User).Find(ctx, bson.M{idName: id})
	if err != nil {
		return nil, err
	}

	var users []model.AppUserDetails

	for cursor.Next(ctx) {
		var result model.AppUserDetails
		err := cursor.Decode(&result)
		if err != nil {
			logs.ErrorLogger.Println("Error retrieving users:", err.Error())
			return nil, err
		}
		users = append(users, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) UpdateUsersStatus(UpdateObject requestDtos.UpdateUserStatus, update primitive.M) (model.AppUserDetails, error) {
	var appUpdateResponse model.AppUserDetails
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

func (r *UserRepository) GetUserEncPW(username string) (model.AppUser, error) {
	var user model.AppUser
	rst, err := connections.GetSessionClient("appusers").Find(context.TODO(), bson.M{"email": username})
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

func (r *UserRepository) GetSingleUserByField(userID string, id string) (model.AppUserDetails, error) {
	var user model.AppUserDetails
	rst, err := connections.GetSessionClient("appusers").Find(context.TODO(), bson.M{id: userID})
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
