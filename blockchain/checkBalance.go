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
	fmt.Println("-----------horizon ", accountRequest)
	account, err := client.AccountDetail(accountRequest)
	if err != nil {
		fmt.Println("---------Not Active 1")
		// log.Fatal(err)
		return "Not active", fmt.Errorf("account does not exist")
	}
	for _, balance := range account.Balances {
		if balance.Asset.Code == "XLM" && balance.Balance >= "10" {
			fmt.Println("balance xlms")
			return balance.Balance, nil
		} else {
			fmt.Println("No Balance")
			return "No Balance", fmt.Errorf("account has a balance of less than 10XLM")
		}
	}
	fmt.Println("Not Active 2")
	return "Not Active", fmt.Errorf("something")
}
