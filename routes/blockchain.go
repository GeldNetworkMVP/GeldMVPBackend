package routes

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/apiModel"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

var BlockchainRoutes = model.Routers{
	//Route will be used to set an account active
	model.Router{
		Name:    "Set New account active",
		Method:  "GET",
		Path:    "/account/activate/{publickey}",
		Handler: apiModel.ActivateNewAccount,
	},
	//Route will be used to fund an account
	model.Router{
		Name:    "Check balance and fund account",
		Method:  "GET",
		Path:    "/account/balance/{publickey}",
		Handler: apiModel.CheckBalance,
	},
}
