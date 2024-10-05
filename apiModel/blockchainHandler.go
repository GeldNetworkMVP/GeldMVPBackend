package apiModel

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/businessFacade"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/errors"
	"github.com/gorilla/mux"
)

func ActivateNewAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	response, err := businessFacade.CheckBalance(vars["publickey"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	}
	if response == "Not active" {
		result, err := businessFacade.ActivateNewAccount(vars["publickey"])
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
	} else {
		fmt.Println("Account is active")
		return
	}
}

func CheckBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := businessFacade.CheckBalance(vars["publickey"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	}
	fmt.Println("-----------1-----------", result)
	if result == "No Balance" {
		response, err := businessFacade.FundAccount(vars["publickey"])
		fmt.Println("----------2-------", response)
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

}
