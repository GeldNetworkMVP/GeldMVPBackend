package routes

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/apiModel"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

var TokenRoutes = model.Routers{
	//Route will be used to add a new template of token data to DB
	model.Router{
		Name:    "Save Token Data",
		Method:  "POST",
		Path:    "/token/save",
		Handler: apiModel.SaveToken,
	},
	//Route will be used to get token by ID
	model.Router{
		Name:    "Get Token Data By ID",
		Method:  "GET",
		Path:    "/token/{_id}",
		Handler: apiModel.GetTokenByID,
	},
	//Route will be used to get all tokens by status paginated call
	model.Router{
		Name:    "Get all tokens by status",
		Method:  "GET",
		Path:    "/tokens/{status}",
		Handler: apiModel.PaginatedGetAllTokensByStatus,
	},
	//Route will be used to update token by ID
	model.Router{
		Name:    "Update token by ID",
		Method:  "PUT",
		Path:    "/tokens/updatestatus",
		Handler: apiModel.UpdateTokenStatus,
	},
}
