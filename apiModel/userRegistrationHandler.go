package apiModel

import (
	"encoding/json"
	"net/http"
	"strconv"

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
func CreateUser(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateUser model.AppUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreateUser)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateUser:userRegistrationHandler: ", err.Error())
	}
	err = validations.ValidateUsers(requestCreateUser)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		encres := commons.Encrypt(requestCreateUser.EncPW)
		requestCreateUser.EncPW = string(encres)
		result, err1 := businessFacade.CreateUsers(requestCreateUser)
		if err1 != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, result)
		}
	}
}

// Trigger the GetUserByID() method that will return The specific User with the ID passed via the API
func GetUsersByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetUserByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		commonResponse.SuccessStatus[model.AppUser](w, result)
	}

}

// Trigger the UpdateUser() method that will return The specific User with the ID passed via the API
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
		commonResponse.SuccessStatus[model.AppUser](w, result)
	}
}

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
**Description:Retrieves all workflows for the specified user ID in a paginated format
 */
func GetPaginatedUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var pagination requestDtos.UserForMatrixView
	pagination.UserID = vars["userid"]
	pgsize, err1 := strconv.Atoi(r.URL.Query().Get("limit"))
	if err1 != nil || pgsize <= 0 {
		_pgsize, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFUALT_LIMIT"))
		logs.InfoLogger.Println("val returned from env: ", _pgsize)
		if envErr != nil {
			errors.InternalError(w, "Something went wrong")
			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
			return
		}
		pgsize = _pgsize
	}
	pagination.PageSize = int32(pgsize)
	requestedPage, err2 := strconv.Atoi(r.URL.Query().Get("page"))
	if err2 != nil || requestedPage <= -1 {
		_requestedpage, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFAULT_PAGE"))
		if envErr != nil {
			errors.InternalError(w, "Something went wrong")
			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
			return
		}
		requestedPage = _requestedpage
	}
	pagination.RequestedPage = int32(requestedPage)
	pagination.SortbyField = "userid"
	sort, err := strconv.Atoi(r.URL.Query().Get("sort"))
	if err != nil || sort != -1 && sort != 1 {
		_sort, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFAULT_PAGE"))
		if envErr != nil {
			errors.InternalError(w, "Something went wrong")
			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
			return
		}
		sort = _sort
	}
	pagination.SortType = sort
	results, err := businessFacade.GetUserDataPagination(pagination)
	if err != nil {
		errors.BadRequest(w, err.Error())
		return
	}
	if results.Content == nil {
		commonResponse.NoContent(w, "")
		return
	}
	commonResponse.SuccessStatus[model.UserPaginatedResponse](w, results)

}
