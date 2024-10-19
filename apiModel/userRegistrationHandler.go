package apiModel

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "strconv"

	"github.com/GeldNetworkMVP/GeldMVPBackend/businessFacade"
	"github.com/GeldNetworkMVP/GeldMVPBackend/commons"
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/commonResponse"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/errors"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/validations"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	All functions here a triggered by api Calls and invokes respective BusinessFace Class Methods
*/
//Retrevies data from the Json Body and decodes it into a model class (User).Then the CreateUser() method is invoked from BusinessFacade
//	@Summary		Create and Save Users in Project Geld
//	@Description	Creates Users for specific designations within the pilot phase of Project Geld
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			usersBody	body		model.UserPayload	true	"User Details"
//	@Success		200			{object}	responseDtos.ResultResponse
//	@Failure		400			{object}	responseDtos.ErrorResponse
//	@Router			/appuser/save [post]
func CreateUser(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateUser model.UserPayload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreateUser)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateUser:userRegistrationHandler: ", err.Error())
	}
	err = validations.ValidateUsers(requestCreateUser)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		encres := commons.Encrypt(requestCreateUser.Password)
		obj := model.AppUser{
			AppUserID: requestCreateUser.AppUserID,
			// AdminUserID: requestCreateUser.AdminUserID,
			Email:       requestCreateUser.Email,
			Contact:     requestCreateUser.Contact,
			Designation: requestCreateUser.Designation,
			EncPW:       encres,
			Status:      requestCreateUser.Status,
			Company:     requestCreateUser.Company,
		}
		result, err1 := businessFacade.CreateUsers(obj)
		if err1 != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, result)
		}
	}
}

// Trigger the GetUserByID() method that will return The specific User with the ID passed via the API
//
//	@Summary		Get user By ID
//	@Description	Get an existing Geld user By ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			_id	path		primitive.ObjectID	true	"AppUserID"
//	@Success		200	{object}	responseDtos.ResultResponse
//	@Failure		400	{object}	responseDtos.ErrorResponse
//	@Router			/appuser/{_id} [get]
func GetUsersByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetUserByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if result.AppUserID == primitive.NilObjectID {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[model.AppUserDetails](w, result)
		}
	}

}

// Trigger the UpdateUser() method that will return The specific User with the ID passed via the API
//
//	@Summary		Update User Details
//	@Description	Update geld user details
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			usersBody	body		requestDtos.UpdateUser	true	"User Details"
//	@Success		200			{object}	responseDtos.ResultResponse
//	@Failure		400			{object}	responseDtos.ErrorResponse
//	@Router			/updateuser [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var UpdateObject requestDtos.UpdateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UpdateObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, err.Error())
		return
	} else {
		result, err := businessFacade.UpdateUsers(UpdateObject)
		if err != nil {
			logs.WarningLogger.Println("Failed to update users : ", err.Error())
			errors.BadRequest(w, err.Error())
			return
		}
		commonResponse.SuccessStatus[model.AppUserDetails](w, result)
	}
}

// DeleteUserByID deletes a user by ID
//
//	@Summary		Delete user By ID
//	@Description	Delete an existing Geld user By ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			_id	path	primitive.ObjectID	true	"AppUserID"
//	@Success		200
//	@Failure		400	{object}	responseDtos.ErrorResponse
//	@Router			/appuser/remove/{_id} [delete]
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	objectID, err := primitive.ObjectIDFromHex(vars["_id"])
	if err != nil {
		logs.WarningLogger.Println("Invalid _id: ", err.Error())
		errors.BadRequest(w, "Invalid _id")
		return
	}
	err1 := businessFacade.DeleteUserByID(objectID)
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		message := "User has been deleted"
		err := json.NewEncoder(w).Encode(message)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
}

/**
**Description:Retrieves all users for the specified adminuser ID in a paginated format
 */
//	@Summary		Get Paginated User Data
//	@Description	Retrieves paginated user data associated with a specific admin user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userid	path		int	true	"UserID"
//	@Param			limit	query		int	false	"Page size (default from env: PAGINATION_DEFUALT_LIMIT)"					Minimum:	1
//	@Param			page	query		int	false	"Requested page (default from env: PAGINATION_DEFAULT_PAGE)"				Minimum:	0
//	@Param			sort	query		int	false	"Sort order (-1: Desc, 1: Asc, default from env: PAGINATION_DEFAULT_SORT)"	Valid		values:	-1,	1
//	@Success		200		{object}	responseDtos.ResultResponse
//	@Success		204		{object}	responseDtos.ResultResponse
//	@Failure		400		{object}	responseDtos.ErrorResponse
//	@Failure		500		{object}	responseDtos.ErrorResponse
//	@Router			/appuser/admin/{userid} [get]

// func GetPaginatedUserData(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json;")
// 	vars := mux.Vars(r)
// 	var pagination requestDtos.UserForMatrixView
// 	pagination.UserID = vars["userid"]
// 	pgsize, err1 := strconv.Atoi(r.URL.Query().Get("limit"))
// 	if err1 != nil || pgsize <= 0 {
// 		_pgsize, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFUALT_LIMIT"))
// 		logs.InfoLogger.Println("val returned from env: ", _pgsize)
// 		if envErr != nil {
// 			errors.InternalError(w, "Something went wrong")
// 			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
// 			return
// 		}
// 		pgsize = _pgsize
// 	}
// 	pagination.PageSize = int32(pgsize)
// 	requestedPage, err2 := strconv.Atoi(r.URL.Query().Get("page"))
// 	if err2 != nil || requestedPage <= -1 {
// 		_requestedpage, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFAULT_PAGE"))
// 		if envErr != nil {
// 			errors.InternalError(w, "Something went wrong")
// 			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
// 			return
// 		}
// 		requestedPage = _requestedpage
// 	}
// 	pagination.RequestedPage = int32(requestedPage)
// 	pagination.SortbyField = "userid"
// 	sort, err := strconv.Atoi(r.URL.Query().Get("sort"))
// 	if err != nil || sort != -1 && sort != 1 {
// 		_sort, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFAULT_PAGE"))
// 		if envErr != nil {
// 			errors.InternalError(w, "Something went wrong")
// 			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
// 			return
// 		}
// 		sort = _sort
// 	}
// 	pagination.SortType = sort
// 	results, err := businessFacade.GetUserDataPagination(pagination)
// 	if err != nil {
// 		errors.BadRequest(w, err.Error())
// 		return
// 	}
// 	if results.Content == nil {
// 		commonResponse.NoContent(w, "")
// 		return
// 	}
// 	commonResponse.SuccessStatus[model.UserPaginatedResponse](w, results)

// }

// Trigger the GetAllUsers() method that will return all the Users
func TestGetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	results, err := businessFacade.TestGetAllUsers()
	if err != nil {
		ErrorMessage := err.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		if len(results) == 0 {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(results)
			if err != nil {
				logs.ErrorLogger.Println("Error occured while encoding JSON in GetAllUsers(UsersHandler): ", err.Error())
			}
			return
		}
	}
}

func GetUsersByStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetUsersByStatus(vars["status"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if len(result) == 0 {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[[]model.AppUserDetails](w, result)
		}
	}

}

func UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var UpdateObject requestDtos.UpdateUserStatus
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UpdateObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, err.Error())
		return
	} else {
		result, err := businessFacade.UpdateUsersStatus(UpdateObject)
		if err != nil {
			logs.WarningLogger.Println("Failed to update users : ", err.Error())
			errors.BadRequest(w, err.Error())
			return
		}
		commonResponse.SuccessStatus[model.AppUserDetails](w, result)
	}
}

func UserSignIn(w http.ResponseWriter, r *http.Request) { //request-body
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var userObj model.AppCredentials
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, err.Error())
		return
	} else {
		result, err := businessFacade.GetUserEncPW(userObj.Email)
		if err != nil {
			logs.WarningLogger.Println("Failed to get users : ", err.Error())
			errors.BadRequest(w, err.Error())
			return
		} else {
			fmt.Println("result: ", result)
			bytepw := result.EncPW
			pw := commons.Decrypt(bytepw)
			fmt.Println("pw---- ", pw)
			if pw == userObj.Pw {
				token, err := commons.GenerateTokenForUser(userObj.Email)
				if err != nil {
					logs.WarningLogger.Println("Failed to generate tokens : ", err.Error())
					errors.BadRequest(w, err.Error())
					return
				} else {
					objectIdString := result.AppUserID.Hex()
					json.NewEncoder(w).Encode(map[string]string{"token": token, "email": result.Email, "contact": result.Contact, "designation": result.Designation, "company": result.Company, "userid": objectIdString})
				}
			} else {
				logs.WarningLogger.Println("Failed validate password")
				errors.BadRequest(w, err.Error())
				return
			}
		}
	}
}

func UserExistence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetUserExistence(vars["email"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if result.AppUserID == primitive.NilObjectID {
			obj := model.UserExistence{
				Status:    "N/A",
				Operative: "No record in existence",
			}
			commonResponse.SuccessStatus[model.UserExistence](w, obj)
		} else {
			fmt.Println("Account exists")
			if result.Status == "accepted" {
				obj := model.UserExistence{
					Status:    "Accepted",
					Operative: result.Designation,
				}
				commonResponse.SuccessStatus[model.UserExistence](w, obj)
			} else if result.Status == "rejected" {
				obj := model.UserExistence{
					Status:    "Rejected",
					Operative: result.Designation,
				}
				commonResponse.SuccessStatus[model.UserExistence](w, obj)
			} else {
				obj := model.UserExistence{
					Status:    "Pending",
					Operative: result.Designation,
				}
				commonResponse.SuccessStatus[model.UserExistence](w, obj)
			}
		}
	}

}
