package businessFacade

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	// "github.com/GeldNetworkMVP/GeldMVPBackend/commons"
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

func SaveTokens(tokens model.Tokens) (string, error) {
	return tokensRepository.SaveTokens(tokens)
}

func SaveTransaction(transaction model.TokenTransactions) (string, error) {
	return tokensRepository.SaveTransactions(transaction)
}

func GetTokenByID(tokensID string) (model.Tokens, error) {
	return tokensRepository.GetTokensByID(tokensID)
}

func UpdateToken(UpdateObject requestDtos.UpdateToken) (model.Tokens, error) {
	update := bson.M{
		"$set": bson.M{"bcstatus": UpdateObject.Status, "tokenhash": UpdateObject.TokenHash},
	}
	return tokensRepository.UpdateTokens(UpdateObject, update)
}

func UpdateTransactions(UpdateObject requestDtos.UpdateToken) (model.TokenTransactions, error) {
	updateTransactions := bson.M{
		"$set": bson.M{"status": UpdateObject.Status, "txnhash": UpdateObject.TokenHash},
	}
	return tokensRepository.UpdateTransactions(UpdateObject, updateTransactions)
}

func GetTokenPaginationByStatus(paginationData requestDtos.TokenForMatrixView) (model.TokenPaginatedresponse, error) {
	filter := bson.M{
		"bcstatus": paginationData.Status,
	}
	projection := GetProjectionDataMatrixViewForTokenData()
	var data []model.Tokens
	response, err := tokensRepository.GetTokensPaginatedResponseByStatus(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "tokens", "bcstatus", data, paginationData.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return model.TokenPaginatedresponse(response), err
	}
	return model.TokenPaginatedresponse(response), err
}

func GetProjectionDataMatrixViewForTokenData() bson.D {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "plotid", Value: 1},
		{Key: "tokenname", Value: 1},
		{Key: "description", Value: 1},
		{Key: "cid", Value: 1},
		{Key: "price", Value: 1},
		{Key: "status", Value: 1},
		{Key: "bcstatus", Value: 1},
		{Key: "tokenhash", Value: 1},
	}
	return projection
}

func GetAllTransactionsByPlotID(plotid string) ([]model.TokenTransactions, error) {
	return tokensRepository.GetAllTransactionsByPlotID("plotid", plotid)
}

func GenerateToken(templates []map[string]interface{}) (string, string, error) {
	return "<html><body><h1>Add some code here</h1></body></html>", "", nil
}

func GetProofBasedOnTemplateTxnHashAndTemplateID(id string, txnhash string) (string, error) {
	var datahash string
	var bchash string
	var dbhash string
	var status string
	url := "https://horizon-testnet.stellar.org/transactions/" + txnhash + "/operations"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logrus.Error("Failed to create request:", err)
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logrus.Error("Failed to execute request:", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Error("Failed to read response body:", err)
		return "", err
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		logrus.Error("Failed to unmarshal response body:", err)
		return "", err
	}

	// Extract embedded operations
	embedded, ok := response["_embedded"].(map[string]interface{})
	if !ok {
		logrus.Error("Missing _embedded field in response")
		return "", fmt.Errorf("missing _embedded field")
	}

	records, ok := embedded["records"].([]interface{})
	if !ok {
		logrus.Error("Missing records field in _embedded response")
		return "", fmt.Errorf("missing records field")
	}

	// Loop through each operation and check if it's a manage_data operation
	for _, record := range records {
		operation, ok := record.(map[string]interface{})
		if !ok {
			continue
		}

		if operation["type"] == "manage_data" {
			entryName := operation["name"].(string)
			if entryName == "Datahash" {
				// Decode base64 value
				value := operation["value"].(string)
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					logrus.Error("Failed to decode base64 value:", err)
					continue
				}

				datahash = string(decodedValue)
				logrus.Info("Decoded manage_data value: ", datahash)

				bchash = datahash
			}
		}
	}

	templateRes, err := dataTemplateRepository.GetTemplateByID(id)
	if err != nil {
		logrus.Error("no document as such in database")
	}

	dbhash = templateRes["templateHash"].(string)
	if dbhash == bchash {
		status = "verified"
	} else {
		status = "error"
	}

	return status, nil

}
