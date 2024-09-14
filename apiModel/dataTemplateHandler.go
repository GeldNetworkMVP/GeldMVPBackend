package apiModel

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/businessFacade"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/commonResponse"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/errors"
	"github.com/gorilla/mux"
)

// Save Data Template submitted
//
//	@Summary		Save Data Templates submitted by the Geld App
//	@Description	The Data Templates aacount for 1 stage which consists of master data and real time data
//	@Tags			data template
//	@Accept			json
//	@Produce		json
//	@Param			dataTemplateBody	body		map[string]interface{}	true	"Data Template Details"
//	@Success		200					{object}	responseDtos.ResultResponse
//	@Failure		400					{object}	responseDtos.ErrorResponse
//	@Router			/geldtemplate/save [post]
func HandlePostTemplateRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	data := map[string]interface{}{}
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Fprintf(w, "Error decoding request: %v", err)
		return
	} else {
		result, err1 := businessFacade.SaveDataTemplate(data)
		if err1 != nil {
			errors.BadRequest(w, "Template with the same name Exists")
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

// Trigger the GetTemplateByID() method that will return The specific template with the ID passed via the API

// @Summary		Get Data Template By ID
// @Description	Get an existing Data Template By ID
// @Tags			data template
// @Accept			json
// @Produce		json
// @Param			_id	path	primitive.ObjectID	true	"TemplateID"
// @Success		200
// @Failure		400	{object}	responseDtos.ErrorResponse
// @Router			/geldtemplate/{_id} [get]
func GetTemplateByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetTemplateByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(result)
		if err != nil {
			fmt.Fprintf(w, "Error encoding response: %v", err)
			return
		}
	}

}

// Trigger the GetTemplateByPlotID() method that will return The specific template with the PlotID passed via the API
//
//	@Summary		Get Data Template By Plot ID
//	@Description	Get an existing Data Template By Plot ID
//	@Tags			data template
//	@Accept			json
//	@Produce		json
//	@Param			plotid	path		string	true	"PlotID"
//	@Success		200		{object}	responseDtos.ResultResponse
//	@Failure		400		{object}	responseDtos.ErrorResponse
//	@Router			/geldtemplate/plotid/{plotid} [get]
func GetTemplateByPlotID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetTemplatesByPlotID(vars["plotid"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if len(result) == 0 {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[[]map[string]interface{}](w, result)
		}
	}

}

// Trigger the GetHTMLByPlotID() method that will return The specific HTML with the PlotID passed via the API
//
//	@Summary		Get HTML By Plot ID
//	@Description	Get an existing HTML for token generation By Plot ID
//	@Tags			data template
//	@Accept			json
//	@Produce		json
//	@Param			plotid	path		string	true	"PlotID"
//	@Success		200		{object}	responseDtos.ResultResponse
//	@Failure		400		{object}	responseDtos.ErrorResponse
//	@Router			/geldtemplate/html/{plotid} [get]
func GetHTMLByPlotID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetWorkflowsByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		commonResponse.SuccessStatus[model.Workflows](w, result)
	}

}

func GetLastTemplateByPlotID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetLastTemplatesByPlotID(vars["plotid"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if len(result) == 0 {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[map[string]interface{}](w, result)
		}
	}

}

func GetTemplateByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetTemplatesByUser(vars["userid"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if len(result) == 0 {
			commonResponse.NotFound(w, "No record found for the given query.")
			return
		} else {
			commonResponse.SuccessStatus[[]map[string]interface{}](w, result)
		}
	}

}
