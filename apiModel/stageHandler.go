package apiModel

import (
	"encoding/json"
	"net/http"

	// "strconv"

	"github.com/GeldNetworkMVP/GeldMVPBackend/businessFacade"
	// "github.com/GeldNetworkMVP/GeldMVPBackend/commons"
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
//Retrevies data from the Json Body and decodes it into a model class (Stage).Then the CreateStage() method is invoked from BusinessFacade
//	@Summary		Create and Save Stages for Workflows
//	@Description	Creates Stages with Fields unique to it for data input labels by Geld App
//	@Tags			stages
//	@Accept			json
//	@Produce		json
//	@Param			stagesBody	body		model.Stages	true	"Stage Details"
//	@Success		200			{object}	responseDtos.ResultResponse
//	@Failure		400			{object}	responseDtos.ErrorResponse
//	@Router			/stage/save [post]
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
		result, err := businessFacade.GetStageExistence(requestCreateStage.StageName)
		if err != nil {
			errors.BadRequest(W, err.Error())
		} else {
			if result.StageName == "" {
				result, err1 := businessFacade.CreateStages(requestCreateStage)
				if err1 != nil {
					errors.BadRequest(W, err.Error())
				} else {
					commonResponse.SuccessStatus[string](W, result)
				}
			} else {
				W.WriteHeader(http.StatusOK)
				message := "Stage with the same name exists"
				err := json.NewEncoder(W).Encode(message)
				if err != nil {
					logs.ErrorLogger.Println(err)
				}
				return
			}
		}
	}
}

// Trigger the GetStageByID() method that will return The specific Stage with the ID passed via the API
//
//	@Summary		Get Stage By ID
//	@Description	Get existing Stage By ID
//	@Tags			stages
//	@Accept			json
//	@Produce		json
//	@Param			_id	path		primitive.ObjectID	true	"StageID"
//	@Success		200	{object}	responseDtos.ResultResponse
//	@Failure		400	{object}	responseDtos.ErrorResponse
//	@Router			/stage/{_id} [get]
func GetStagesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetStagesByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if result.StageID == primitive.NilObjectID {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[model.Stages](w, result)
		}
	}

}

// Trigger the UpdateStages() method that will return The specific updated Stage with the ID passed via the API
// Trigger the UpdateMasterData() method that will return The specific updated MasterData with the ID passed via the API
//
//	@Summary		Update Stage Data
//	@Description	Update an exisiting Stages
//	@Tags			stages
//	@Accept			json
//	@Produce		json
//	@Param			stagesBody	body		requestDtos.UpdateStages	true	"Stage Details"
//	@Success		200			{object}	responseDtos.ResultResponse
//	@Failure		400			{object}	responseDtos.ErrorResponse
//	@Router			/updatestage [put]
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

// DeleteStageByID deletes a stage by ID
//
//	@Summary		Delete Stage By ID
//	@Description	Delete an existing Stage By ID
//	@Tags			stages
//	@Accept			json
//	@Produce		json
//	@Param			_id	path	primitive.ObjectID	true	"StageID"
//	@Success		200
//	@Failure		400	{object}	responseDtos.ErrorResponse
//	@Router			/stage/remove/{_id} [delete]
func DeleteStageByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	objectID, err := primitive.ObjectIDFromHex(vars["_id"])
	if err != nil {
		logs.WarningLogger.Println("Invalid _id: ", err.Error())
		errors.BadRequest(w, "Invalid _id")
		return
	}
	err1 := businessFacade.DeleteStageByID(objectID)
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
//	@Summary		Get Paginated Stage Data
//	@Description	Retrieves paginated Stage data associated with a specific user
//	@Tags			stages
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
//	@Router			/userstage/{userid} [get]

// func GetPaginatedStageData(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json;")
// 	vars := mux.Vars(r)
// 	var pagination requestDtos.StagesForMatrixView
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
// 	results, err := businessFacade.GetStageDataPagination(pagination)
// 	if err != nil {
// 		errors.BadRequest(w, err.Error())
// 		return
// 	}
// 	if results.Content == nil {
// 		commonResponse.NoContent(w, "")
// 		return
// 	}
// 	commonResponse.SuccessStatus[model.StagePaginatedresponse](w, results)

// }

func GetStagesByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetStagesByName(vars["stagename"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if result.StageID == primitive.NilObjectID {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[model.Stages](w, result)
		}
	}

}

// Trigger the GetAllStages() method that will return all the Stages
func TestGetAllStages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	results, err := businessFacade.TestGetAllStages()
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
				logs.ErrorLogger.Println("Error occured while encoding JSON in GetAllStages(StagesHandler): ", err.Error())
			}
			return
		}
	}
}

func FilterExistingStages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var stages model.StagesNames
	var validStages []string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&stages)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateStages:stagesHandler: ", err.Error())
	} else {
		for _, stage := range stages.StageArray {
			result, err := businessFacade.GetStageExistence(stage)
			if err != nil {
				errors.BadRequest(w, err.Error())
			} else {
				if result.StageName != "" {
					validStages = append(validStages, stage)
				}
			}
		}
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(validStages)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}

}
