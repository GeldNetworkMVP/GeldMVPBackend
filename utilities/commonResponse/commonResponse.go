package commonResponse

import (
	"encoding/json"
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/responseDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
)

type resultType interface {
	model.Workflows | string | model.MasterData | model.Stages | model.DataCollection | []model.DataCollection | []model.Workflows | []model.Stages | []model.MasterData | model.MDataPaginatedresponse | model.DataPaginatedresponse | model.StagePaginatedresponse | model.WorkflowPaginatedresponse | model.AppUser | model.UserPaginatedResponse | model.Tokens | model.TokenPaginatedresponse | []map[string]interface{} | map[string]interface{} | []model.AppUser | []model.TokenTransactions | model.AppUserDetails | []model.AppUserDetails | model.UserExistence | model.TokenPostResult
}

func SuccessStatus[T resultType](w http.ResponseWriter, result T) {
	w.WriteHeader(http.StatusOK)
	response := responseDtos.ResultResponse{
		Status:   http.StatusOK,
		Response: result,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}

func NoContent(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNoContent)
	response := responseDtos.ErrorResponse{
		Message: message,
		Status:  http.StatusNoContent,
		Error:   "No record in result",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}

func RespondWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func NotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	response := responseDtos.ErrorResponse{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "No record found in result",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}
