package apiModel

import (
	"encoding/json"
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/businessFacade"
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
//Retrevies data from the Json Body and decodes it into a model class (Master Data).Then the CreateMasterData() method is invoked from BusinessFacade
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

// Trigger the UpdateMasterData() method that will return The specific updated MasterData with the ID passed via the API
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
