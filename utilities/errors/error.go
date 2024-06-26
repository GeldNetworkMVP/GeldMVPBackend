package errors

import (
	"encoding/json"
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/responseDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
)

func BadRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	response := responseDtos.ErrorResponse{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad request",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}

func NotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	response := responseDtos.ErrorResponse{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not found",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}

func InternalError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	response := responseDtos.ErrorResponse{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal server error",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}

func DBCoonectionIssue(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	response := responseDtos.ErrorResponse{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Databse connection issue",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}
