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
)

/*
	All functions here a triggered by api Calls and invokes respective BusinessFace Class Methods
*/
//Retrevies data from the Json Body and decodes it into a model class (Stage).Then the CreateStage() method is invoked from BusinessFacade
func CreateStage(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateStage model.Stages
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreateStage)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateStages:stagesHandler: ", err.Error())
	}
	err = validations.ValidateStages(requestCreateStage)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		result, err1 := businessFacade.CreateStages(requestCreateStage)
		if err1 != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, result)
		}
	}
}

// Trigger the GetStageByID() method that will return The specific Stage with the ID passed via the API
func GetStagesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetStagesByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		commonResponse.SuccessStatus[model.Stages](w, result)
	}

}

// Trigger the UpdateStages() method that will return The specific updated Stage with the ID passed via the API
func UpdateStages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var UpdateObject requestDtos.UpdateStages
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UpdateObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, err.Error())
		return
	} else {
		result, err := businessFacade.UpdateStages(UpdateObject)
		if err != nil {
			logs.WarningLogger.Println("Failed to update stage : ", err.Error())
			errors.BadRequest(w, err.Error())
			return
		}
		commonResponse.SuccessStatus[model.Stages](w, result)
	}
}

func DeleteStageByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)

	err1 := businessFacade.DeleteStageByID(vars["_id"])
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		message := "Stage has been deleted"
		err := json.NewEncoder(w).Encode(message)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
}

/**
**Description:Retrieves all stages for the specified user ID in a paginated format
 */
func GetPaginatedStageData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var pagination requestDtos.StagesForMatrixView
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
	pagination.SortbyFeild = "userid"
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
	results, err := businessFacade.GetStageDataPagination(pagination)
	if err != nil {
		errors.BadRequest(w, err.Error())
		return
	}
	if results.Content == nil {
		commonResponse.NoContent(w, "")
		return
	}
	commonResponse.SuccessStatus[model.StagePaginatedresponse](w, results)

}
