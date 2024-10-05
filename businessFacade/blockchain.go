package businessFacade

import "github.com/GeldNetworkMVP/GeldMVPBackend/blockchain"

func ActivateNewAccount(publickey string) (string, error) {
	return blockchain.ActivateAccount(publickey)
}

func FundAccount(publickey string) (string, error) {
	return blockchain.FundAccount(publickey)
}

func CheckBalance(publickey string) (string, error) {
	return blockchain.CheckBalance(publickey)
}
