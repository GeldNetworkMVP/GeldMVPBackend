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
	//Route will be used to get data template according to PlotID
	model.Router{
		Name:    "Get DataTemplate By PlotID",
		Method:  "GET",
		Path:    "/geldtemplate/plotid/{plotid}",
		Handler: apiModel.GetTemplateByPlotID,
	},
	//Route will be used to get data template according to ID
	model.Router{
		Name:    "Get DataTemplate By ID",
		Method:  "GET",
		Path:    "/geldtemplate/{_id}",
		Handler: apiModel.GetTemplateByID,
	},
	//Route will be used to get HTML according to PlotID
	model.Router{
		Name:    "Get HTML By PlotID",
		Method:  "GET",
		Path:    "/geldtemplate/html/{plotid}",
		Handler: apiModel.GetHTMLByPlotID,
	},
	//Route will be used to get last data template according to PlotID
	model.Router{
		Name:    "Get Last DataTemplate By PlotID",
		Method:  "GET",
		Path:    "/lastgeldtemplate/plotid/{plotid}",
		Handler: apiModel.GetLastTemplateByPlotID,
	},
	model.Router{
		Name:    "Get DataTemplate By User",
		Method:  "GET",
		Path:    "/geldtemplate/user/{userid}",
		Handler: apiModel.GetTemplateByUser,
	},
}
