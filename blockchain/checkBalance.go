package blockchain

import (
	"fmt"
	// "log"

	"github.com/GeldNetworkMVP/GeldMVPBackend/commons"
	"github.com/stellar/go/clients/horizonclient"
)

func CheckBalance(publickey string) (string, error) {
	client := commons.GetHorizonClient()
	accountRequest := horizonclient.AccountRequest{AccountID: publickey}
	account, err := client.AccountDetail(accountRequest)
	if err != nil {
		return "Not active", nil
	} else {
		for _, balance := range account.Balances {
			if balance.Asset.Type == "native" && balance.Balance >= "10" {
				return "Funded", nil
			} else {
				return "No Balance", nil
			}
		}
	}
	return "Not active", fmt.Errorf("something")
}
