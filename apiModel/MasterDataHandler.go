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
//Retrevies data from the Json Body and decodes it into a model class (Master Data).Then the CreateMasterData() method is invoked from BusinessFacade

// @Summary		Save Master Data Container submitted by Geld Configurations
// @Description	This creates a container to whole master data records per type
// @Tags			master data container
// @Accept			json
// @Produce		json
// @Param			masterDataContainerBody	body		model.MasterData	true	"Master Data Container Details"
// @Success		200						{object}	responseDtos.ResultResponse
// @Failure		400						{object}	responseDtos.ErrorResponse
// @Router			/masterdata/save [post]
func CreateMasterData(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateMdata model.MasterData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreateMdata)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateMasterData:masterDataHandler: ", err.Error())
	}
	err = validations.ValidateMasterData(requestCreateMdata)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		result, err1 := businessFacade.CreateMasterData(requestCreateMdata)
		if err1 != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, result)
		}
	}
}

// Retrevies data from the Json Body and decodes it into a model class (Data Collection).Then the CreateDataCollection() method is invoked from BusinessFacade

// @Summary		Save Data Collection for Master Data Container
// @Description	This creates a collections for Master Data Containers based on container type
// @Tags			data collection
// @Accept			json
// @Produce		json
// @Param			dataCollectionBody	body		model.DataCollection	true	"Data Collection Details"
// @Success		200					{object}	responseDtos.ResultResponse
// @Failure		400					{object}	responseDtos.ErrorResponse
// @Router			/record/save [post]
func CreateDataCollection(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateMdataCollection model.DataCollection
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreateMdataCollection)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateMasterDataCollection:masterDataHandler: ", err.Error())
	}
	err = validations.ValidateMasterDataCollection(requestCreateMdataCollection)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		result, err1 := businessFacade.CreateMasterDataCollection(requestCreateMdataCollection)
		if err1 != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, result)
		}
	}
}

// Trigger the GetMasterDataByID() method that will return The specific MasterData with the ID passed via the API
//
//	@Summary		Get Master Data by ID
//	@Description	Get an existing Master Data Container ID
//	@Tags			master data container
//	@Accept			json
//	@Produce		json
//	@Param			_id	path		primitive.ObjectID	true	"DataID"
//	@Success		200	{object}	responseDtos.ResultResponse
//	@Failure		400	{object}	responseDtos.ErrorResponse
//	@Router			/masterdata/{_id} [get]
func GetMasterDataByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetMasterDataByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		commonResponse.SuccessStatus[model.MasterData](w, result)
	}

}

// Trigger the GetRecordDataByID() method that will return The specific RecordData with the ID passed via the API
//
//	@Summary		Get Record Data By ID
//	@Description	Get an existing Record Data By ID
//	@Tags			data collection
//	@Accept			json
//	@Produce		json
//	@Param			_id	path		primitive.ObjectID	true	"CollectionID"
//	@Success		200	{object}	responseDtos.ResultResponse
//	@Failure		400	{object}	responseDtos.ErrorResponse
//	@Router			/record/{_id} [get]
func GetRecordDataByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetRecordDataByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if result.CollectionID == primitive.NilObjectID {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[model.DataCollection](w, result)
		}
	}

}

// Trigger the GetRecordDataByMasterDataID() method that will return The specific RecordData with the MasterDataID passed via the API
//
//	@Summary		Get Record Data By Master Data ID
//	@Description	Get an existing Record Data By Master Data ID
//	@Tags			data collection
//	@Accept			json
//	@Produce		json
//	@Param			dataid	path		string	true	"DataID"
//	@Success		200		{object}	responseDtos.ResultResponse
//	@Failure		400		{object}	responseDtos.ErrorResponse
//	@Router			/records/{dataid} [get]
func GetRecordDataByMasterDataID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetRecordDataByMasterDataID(vars["dataid"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if len(result) == 0 {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[[]model.DataCollection](w, result)
		}
	}

}

/**
**Description:Retrieves all masterdata for the specified user in a paginated format
 */
//	@Summary		Get Paginated master data
//	@Description	Retrieves paginated master data associated with a specific user
//	@Tags			master data container
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
//	@Router			/usermasterdata/{userid} [get]

// func GetPaginatedMasterData(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json;")
// 	vars := mux.Vars(r)
// 	var pagination requestDtos.MasterDataForMatrixView
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
// 	results, err := businessFacade.GetMasterDataPagination(pagination)
// 	if err != nil {
// 		errors.BadRequest(w, err.Error())
// 		return
// 	}
// 	if results.Content == nil {
// 		commonResponse.NoContent(w, "")
// 		return
// 	}
// 	commonResponse.SuccessStatus[model.MDataPaginatedresponse](w, results)

// }

/**
**Description:Retrieves all data collections for the specified master data ID in a paginated format
 */
//	@Summary		Get Paginated data collection
//	@Description	Retrieves paginated data collection associated with a specific dataID
//	@Tags			data collection
//	@Accept			json
//	@Produce		json
//	@Param			dataid	path		int	true	"DataID"
//	@Param			limit	query		int	false	"Page size (default from env: PAGINATION_DEFUALT_LIMIT)"					Minimum:	1
//	@Param			page	query		int	false	"Requested page (default from env: PAGINATION_DEFAULT_PAGE)"				Minimum:	0
//	@Param			sort	query		int	false	"Sort order (-1: Desc, 1: Asc, default from env: PAGINATION_DEFAULT_SORT)"	Valid		values:	-1,	1
//	@Success		200		{object}	responseDtos.ResultResponse
//	@Success		204		{object}	responseDtos.ResultResponse
//	@Failure		400		{object}	responseDtos.ErrorResponse
//	@Failure		500		{object}	responseDtos.ErrorResponse
//	@Router			/masterrecord/{dataid} [get]
func GetPaginatedData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var pagination requestDtos.DataRecordForMatrixView
	pagination.DataID = vars["dataid"]
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
	pagination.SortbyField = "dataid"
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
	results, err := businessFacade.GetDataPagination(pagination)
	if err != nil {
		errors.BadRequest(w, err.Error())
		return
	}
	if results.Content == nil {
		commonResponse.NoContent(w, "")
		return
	}
	commonResponse.SuccessStatus[model.DataPaginatedresponse](w, results)

}

// Trigger the UpdateMasterData() method that will return The specific updated MasterData with the ID passed via the API
//
//	@Summary		Update Master Data Container
//	@Description	Update an exisiting Master Data Container
//	@Tags			master data container
//	@Accept			json
//	@Produce		json
//	@Param			masterDataContainerBody	body		requestDtos.UpdateMasterData	true	"Master Data Container Details"
//	@Success		200						{object}	responseDtos.ResultResponse
//	@Failure		400						{object}	responseDtos.ErrorResponse
//	@Router			/updatemasterdata [put]
func UpdateMasterData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var UpdateObject requestDtos.UpdateMasterData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UpdateObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, err.Error())
		return
	} else {
		result, err := businessFacade.UpdateMasterData(UpdateObject)
		if err != nil {
			logs.WarningLogger.Println("Failed to update masterdata : ", err.Error())
			errors.BadRequest(w, err.Error())
			return
		}
		commonResponse.SuccessStatus[model.MasterData](w, result)
	}
}

// Trigger the UpdateDataCollection() method that will return The specific updated MasterDataollection with the ID passed via the API
//
//	@Summary		Update Master Data Container
//	@Description	Update an exisiting Master Data Container
//	@Tags			data collection
//	@Accept			json
//	@Produce		json
//	@Param			dataCollectionBody	body		requestDtos.UpdateDataCollection	true	"Data Collection Details"
//	@Success		200					{object}	responseDtos.ResultResponse
//	@Failure		400					{object}	responseDtos.ErrorResponse
//	@Router			/updaterecords [put]
func UpdateDataCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var UpdateObject requestDtos.UpdateDataCollection
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UpdateObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, err.Error())
		return
	} else {
		result, err := businessFacade.UpdateDataCollection(UpdateObject)
		if err != nil {
			logs.WarningLogger.Println("Failed to update masterdata : ", err.Error())
			errors.BadRequest(w, err.Error())
			return
		}
		commonResponse.SuccessStatus[model.DataCollection](w, result)
	}
}

// DeleteMasterDataByID deletes an master data by ID
//
//	@Summary		Delete Master Data Container By ID
//	@Description	Delete an existing Master Data Container By ID
//	@Tags			master data container
//	@Accept			json
//	@Produce		json
//	@Param			_id	path	primitive.ObjectID	true	"DataID"
//	@Success		200
//	@Failure		400	{object}	responseDtos.ErrorResponse
//	@Router			/masterdata/remove/{_id} [delete]
func DeleteMasterDataByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	objectID, err := primitive.ObjectIDFromHex(vars["_id"])
	if err != nil {
		logs.WarningLogger.Println("Invalid _id: ", err.Error())
		errors.BadRequest(w, "Invalid _id")
		return
	}
	err1 := businessFacade.DeleteMasterDataByID(objectID)
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		message := "MasterData has been deleted"
		err := json.NewEncoder(w).Encode(message)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
}

// DeleteMasterDataRecordByID deletes an master data records by ID
//
//	@Summary		Delete Master Data Records By ID
//	@Description	Delete an existing Master Data Record By ID
//	@Tags			data collection
//	@Accept			json
//	@Produce		json
//	@Param			_id	path	primitive.ObjectID	true	"CollectionID"
//	@Success		200
//	@Failure		400	{object}	responseDtos.ErrorResponse
//	@Router			/record/remove/{_id} [delete]
func DeleteMasterDataRecordByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	objectID, err := primitive.ObjectIDFromHex(vars["_id"])
	if err != nil {
		logs.WarningLogger.Println("Invalid _id: ", err.Error())
		errors.BadRequest(w, "Invalid _id")
		return
	}
	err1 := businessFacade.DeleteMasterDataRecordByID(objectID)
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		message := "Record has been deleted"
		err := json.NewEncoder(w).Encode(message)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
}

func GetPlotDataByMasterDataID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	result, err := businessFacade.GetRecordDataByMasterDataID("66d71f5af2999f18ef9f9d30")
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if len(result) == 0 {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[[]model.DataCollection](w, result)
		}
	}

}
