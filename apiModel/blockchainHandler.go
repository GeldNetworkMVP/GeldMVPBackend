package apiModel

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/blockchain/tokens"
	"github.com/GeldNetworkMVP/GeldMVPBackend/businessFacade"
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/errors"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"github.com/gorilla/mux"
)

func ActivateNewAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var UpdateObject requestDtos.UpdateUserPublicKey
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UpdateObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, err.Error())
		return
	} else {
		response, err := businessFacade.CheckBalance(UpdateObject.PublicKey)
		if err != nil {
			errors.BadRequest(w, err.Error())
		}
		if response == "Not active" {
			result, err := businessFacade.ActivateNewAccount(UpdateObject.PublicKey)
			if err != nil {
				errors.BadRequest(w, err.Error())
				return
			} else {
				res, err := businessFacade.UpdateUsersPublicKey(UpdateObject)
				if err != nil && res.PublicKey == "" {
					errors.BadRequest(w, err.Error())
					return
				} else {
					w.Header().Set("Content-Type", "application/json")
					err := json.NewEncoder(w).Encode(result)
					if err != nil {
						fmt.Fprintf(w, "Error encoding response: %v", err)
						return
					}
				}
			}
		} else {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode("Account is active")
			if err != nil {
				fmt.Fprintf(w, "Error encoding response: %v", err)
				return
			}
		}
	}
}

func CheckBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.CheckBalance(vars["publickey"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	}
	if result == "No Balance" {
		response, err := businessFacade.FundAccount(vars["publickey"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(response)
			if err != nil {
				fmt.Fprintf(w, "Error encoding response: %v", err)
				return
			}
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(result)
		if err != nil {
			fmt.Fprintf(w, "Error encoding response: %v", err)
			return
		}
	}

}

func GetIssuerAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var TokenIssuerPK, EncodedTokenIssuerSK, err = tokens.CreateIssuer()
	if err != nil && TokenIssuerPK == "" && EncodedTokenIssuerSK == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		var TokenKeys = model.Keys{
			PK: TokenIssuerPK,
			SK: EncodedTokenIssuerSK,
		}
		res, err := businessFacade.SaveStellarKeys(TokenKeys)
		if err != nil {
			fmt.Println(err)
			return
		} else {

		}
		//send the response
		result := model.IssuerResponse{
			IssuerPK: TokenIssuerPK,
			Result:   res,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
