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
//Retrevies data from the Json Body and decodes it into a model class (Workflow).Then the CreateWorkflow() method is invoked from BusinessFacade
//	@Summary		Create and Save Workflows for Plots
//	@Description	Creates workflows to have a series of stages to map plot chain management
//	@Tags			workflows
//	@Accept			json
//	@Produce		json
//	@Param			workflowsBody	body		model.Workflows	true	"Workflow Details"
//	@Success		200				{object}	responseDtos.ResultResponse
//	@Failure		400				{object}	responseDtos.ErrorResponse
//	@Router			/workflows/save [post]
func CreateWorkflow(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateWorkflow model.Workflows
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreateWorkflow)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateWorkflow:workflowHandler: ", err.Error())
	}
	err = validations.ValidateWorkflows(requestCreateWorkflow)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		result, err := businessFacade.GetWorkflowExistence(requestCreateWorkflow.WorkflowName)
		if err != nil {
			errors.BadRequest(W, err.Error())
		} else {
			if result.WorkflowName == "" {
				result, err1 := businessFacade.CreateWorkflows(requestCreateWorkflow)
				if err1 != nil {
					errors.BadRequest(W, err.Error())
				} else {
					commonResponse.SuccessStatus[string](W, result)
				}
			} else {
				W.WriteHeader(http.StatusOK)
				message := "Workflow with the same name exists"
				err := json.NewEncoder(W).Encode(message)
				if err != nil {
					logs.ErrorLogger.Println(err)
				}
				return
			}
		}

	}
}

// Trigger the GetWorkflowByID() method that will return The specific Workflow with the ID passed via the API
//
//	@Summary		Get workflow By ID
//	@Description	Get an existing workflow By ID
//	@Tags			workflows
//	@Accept			json
//	@Produce		json
//	@Param			_id	path		primitive.ObjectID	true	"WorkflowID"
//	@Success		200	{object}	responseDtos.ResultResponse
//	@Failure		400	{object}	responseDtos.ErrorResponse
//	@Router			/workflows/{_id} [get]
func GetWorkflowsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetWorkflowsByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if result.WorkflowID == primitive.NilObjectID {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[model.Workflows](w, result)
		}
	}

}

// Trigger the UpdateWorkflow() method that will return The specific Workflow with the ID passed via the API
//
//	@Summary		Update Workflow Details
//	@Description	Update Workflow details
//	@Tags			workflows
//	@Accept			json
//	@Produce		json
//	@Param			workflowsBody	body		requestDtos.UpdateWorkflow	true	"Workflow Details"
//	@Success		200				{object}	responseDtos.ResultResponse
//	@Failure		400				{object}	responseDtos.ErrorResponse
//	@Router			/updateworkflow [put]
func UpdateWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var UpdateObject requestDtos.UpdateWorkflow
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UpdateObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, err.Error())
		return
	} else {
		result, err := businessFacade.UpdateWorkflow(UpdateObject)
		if err != nil {
			logs.WarningLogger.Println("Failed to update workflow : ", err.Error())
			errors.BadRequest(w, err.Error())
			return
		}
		commonResponse.SuccessStatus[model.Workflows](w, result)
	}
}

// DeleteWorkflowByID deletes a workflow by ID
//
//	@Summary		Delete workflow By ID
//	@Description	Delete an existing workflow By ID
//	@Tags			workflows
//	@Accept			json
//	@Produce		json
//	@Param			_id	path	primitive.ObjectID	true	"WorkflowID"
//	@Success		200
//	@Failure		400	{object}	responseDtos.ErrorResponse
//	@Router			/workflows/remove/{_id} [delete]
func DeleteWorkflowByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	objectID, err := primitive.ObjectIDFromHex(vars["_id"])
	if err != nil {
		logs.WarningLogger.Println("Invalid _id: ", err.Error())
		errors.BadRequest(w, "Invalid _id")
		return
	}
	err1 := businessFacade.DeleteWorkflowByID(objectID)
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		message := "Workflow has been deleted"
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
//	@Summary		Get Paginated Workflow Data
//	@Description	Retrieves paginated workflow data associated with a specific user
//	@Tags			workflows
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
//	@Router			/userworkflows/{userid} [get]

//

func TestPaginatedWorkflowData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var pagination requestDtos.WorkflowForMatrixView

	// Extract query parameters for pagination
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

	// Remove sortbyField setting since we're not sorting by UserID anymore
	sort, err := strconv.Atoi(r.URL.Query().Get("sort"))
	if err != nil || (sort != -1 && sort != 1) {
		_sort, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFAULT_PAGE"))
		if envErr != nil {
			errors.InternalError(w, "Something went wrong")
			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
			return
		}
		sort = _sort
	}
	pagination.SortType = sort

	// Call the business logic to get paginated workflow data
	results, err := businessFacade.TestWorkflowDataPagination(pagination)
	if err != nil {
		errors.BadRequest(w, err.Error())
		return
	}

	if results.Content == nil {
		commonResponse.NoContent(w, "")
		return
	}

	// Return the results in JSON format
	commonResponse.SuccessStatus[model.WorkflowPaginatedresponse](w, results)
}

// Trigger the GetAllWorkflows() method that will return all the Workflows
func TestGetAllWorkflows(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	results, err := businessFacade.TestGetAllWorkflows()
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
				logs.ErrorLogger.Println("Error occured while encoding JSON in GetAllWorkflow(WorkflowHandler): ", err.Error())
			}
			return
		}
	}
}
