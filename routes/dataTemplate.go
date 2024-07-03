package routes

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/apiModel"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

var DataTemplateRoutes = model.Routers{
	//Route will be used to add a new template of data to DB
	model.Router{
		Name:    "Save DataTemplate",
		Method:  "POST",
		Path:    "/geldtemplate/save",
		Handler: apiModel.HandlePostTemplateRequest,
	},
}
