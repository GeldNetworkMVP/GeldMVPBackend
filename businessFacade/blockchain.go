package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/blockchain"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

func ActivateNewAccount(publickey string) (string, error) {
	return blockchain.ActivateAccount(publickey)
}

func FundAccount(publickey string) (string, error) {
	return blockchain.FundAccount(publickey)
}

func CheckBalance(publickey string) (string, error) {
	return blockchain.CheckBalance(publickey)
}

func SaveStellarKeys(keys model.Keys) (string, error) {
	return tokensRepository.SaveStellarKeys(keys)
}

func GetStellarKeys(pk string) (model.Keys, error) {
	return tokensRepository.GetStellarKeys(pk)
}
