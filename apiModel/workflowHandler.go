package apiModel

import (
	"encoding/json"
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/businessFacade"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/commonResponse"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/errors"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/validations"
)

/*
	All functions here a triggered by api Calls and invokes respective BusinessFace Class Methods
*/
//Retrevies data from the Json Body and decodes it into a model class (Workflow).Then the CreateWorkflow() method is invoked from BusinessFacade
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
		result, err1 := businessFacade.CreateWorkflows(requestCreateWorkflow)
		if err1 != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, result)
		}
	}
}
