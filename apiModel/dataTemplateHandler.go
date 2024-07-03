package apiModel

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GeldNetworkMVP/GeldMVPBackend/businessFacade"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/commonResponse"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/errors"
)

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
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}
