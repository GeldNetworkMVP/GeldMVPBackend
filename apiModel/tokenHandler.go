package apiModel

import (
	"encoding/json"
	"fmt"
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
//Retrevies data from the Json Body and decodes it into a model class (Token).Then the SaveToken() method is invoked from BusinessFacade
//	@Summary		Create and Save Tokens per Plot
//	@Description	Gets a timeline of data templates and creates a token per plot
//	@Tags			tokens
//	@Accept			json
//	@Produce		json
//	@Param			tokensBody	body		model.TokenPayload	true	"Token Details"
//	@Success		200			{object}	responseDtos.ResultResponse
//	@Failure		400			{object}	responseDtos.ErrorResponse
//	@Router			/token/save [post]
func SaveToken(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateToken model.TokenPayload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreateToken)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in saveToken:tokenHandler: ", err.Error())
	}
	err = validations.ValidateToken(requestCreateToken)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		templates, err := businessFacade.GetTemplatesByPlotID(requestCreateToken.PlotID)
		if err != nil {
			errors.BadRequest(W, err.Error())
		} else {
			svg, tokenhash, err := businessFacade.GenerateToken(templates)
			if err != nil {
				errors.BadRequest(W, err.Error())
			} else {
				cid, err := businessFacade.UploadFilesToIpfs(requestCreateToken, svg)
				if err != nil {
					errors.BadRequest(W, err.Error())
				} else {
					object := model.Tokens{
						TokenID:     requestCreateToken.TokenID,
						PlotID:      requestCreateToken.PlotID,
						TokenName:   requestCreateToken.TokenName,
						Description: requestCreateToken.Description,
						CID:         cid,
						Price:       requestCreateToken.Price,
						IPFSStatus:  "Sent to IPFS",
						BCStatus:    requestCreateToken.BCStatus,
						TokenHash:   tokenhash,
					}
					result, err1 := businessFacade.SaveTokens(object)
					if err1 != nil {
						errors.BadRequest(W, err.Error())
					} else {
						//TODO: check BC Status as well and do a comparison later on
						newobj := model.TokenTransactions{
							TransactionStatus: "OnSale",
							TokenName:         requestCreateToken.TokenName,
							TXNHash:           requestCreateToken.BCHash,
							PlotID:            requestCreateToken.PlotID,
							TokenID:           result,
							DBStatus:          "Saved",
						}
						result1, err2 := businessFacade.SaveTransaction(newobj)
						if err2 != nil {
							errors.BadRequest(W, err.Error())
						} else {
							commonResponse.SuccessStatus[string](W, result1)
						}
					}
				}
			}
		}
	}
}

// Trigger the GetTokenByID() method that will return The specific Token with the ID passed via the API
//
//	@Summary		Get Token By ID
//	@Description	Get Existing Token By ID
//	@Tags			tokens
//	@Accept			json
//	@Produce		json
//	@Param			_id	path	primitive.ObjectID	true	"TokenID"
//	@Success		200	{array}	responseDtos.ResultResponse
//	@Failure		400	{array}	responseDtos.ErrorResponse
//	@Router			/token/{_id} [get]
func GetTokenByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetTokenByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		commonResponse.SuccessStatus[model.Tokens](w, result)
	}

}

// Trigger the UpdateTokenStatus() method that will return The specific Token with the ID passed via the API
//
//	@Summary		Update Token Status
//	@Description	Update Token Status based on if it has been minted or sold and bought
//	@Tags			tokens
//	@Accept			json
//	@Produce		json
//	@Param			tokensBody	body		requestDtos.UpdateToken	true	"Token Details"
//	@Success		200			{object}	responseDtos.ResultResponse
//	@Failure		400			{object}	responseDtos.ErrorResponse
//	@Router			/tokens/updatestatus [put]
func UpdateTokenStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var UpdateObject requestDtos.UpdateToken
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UpdateObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, err.Error())
		return
	} else {
		result, err := businessFacade.UpdateToken(UpdateObject)
		if err != nil {
			logs.WarningLogger.Println("Failed to update token : ", err.Error())
			errors.BadRequest(w, err.Error())
			return
		} else {
			resultTxn, err := businessFacade.UpdateTransactions(UpdateObject)
			logs.InfoLogger.Println("Transaction Data has been updated ", resultTxn)
			if err != nil {
				logs.WarningLogger.Println("Failed to update token : ", err.Error())
				errors.BadRequest(w, err.Error())
				return
			}
			commonResponse.SuccessStatus[model.Tokens](w, result)
		}

	}
}

//	@Summary		Get Paginated Token Data
//	@Description	Retrieves paginated token data associated with a specific status
//	@Tags			tokens
//	@Accept			json
//	@Produce		json
//	@Param			status	path		int	true	"Status"
//	@Param			limit	query		int	false	"Page size (default from env: PAGINATION_DEFUALT_LIMIT)"					Minimum:	1
//	@Param			page	query		int	false	"Requested page (default from env: PAGINATION_DEFAULT_PAGE)"				Minimum:	0
//	@Param			sort	query		int	false	"Sort order (-1: Desc, 1: Asc, default from env: PAGINATION_DEFAULT_SORT)"	Valid		values:	-1,	1
//	@Success		200		{object}	responseDtos.ResultResponse
//	@Success		204		{object}	responseDtos.ResultResponse
//	@Failure		400		{object}	responseDtos.ErrorResponse
//	@Failure		500		{object}	responseDtos.ErrorResponse
//	@Router			/tokens/{status} [get]

func PaginatedGetAllTokensByStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var pagination requestDtos.TokenForMatrixView
	pagination.Status = vars["status"]
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
	pagination.SortbyField = "status"
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
	results, err := businessFacade.GetTokenPaginationByStatus(pagination)
	if err != nil {
		errors.BadRequest(w, err.Error())
		return
	}
	if results.Content == nil {
		commonResponse.NoContent(w, "")
		return
	}
	commonResponse.SuccessStatus[model.TokenPaginatedresponse](w, results)

}

func GetAllTransactionsByPlotID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetAllTransactionsByPlotID(vars["plotid"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if len(result) == 0 {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(result)
			if err != nil {
				fmt.Fprintf(w, "Error encoding response: %v", err)
				return
			}
		} else {
			commonResponse.SuccessStatus[[]model.TokenTransactions](w, result)
		}
	}

}

func GetProofBasedOnTemplateTxnHashAndTemplateID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.GetProofBasedOnTemplateTxnHashAndTemplateID(vars["_id"], vars["currentHash"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		// if len(result) == 0 {
		// 	commonResponse.NotFound(w, "No record found for the given query.")
		// 	return
		// } else {
		commonResponse.SuccessStatus[string](w, result)
		//}
	}

}
