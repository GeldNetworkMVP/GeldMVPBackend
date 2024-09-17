package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveDataTemplate(model map[string]interface{}) (string, error) {
	return dataTemplateRepository.SaveDataTemplate(model)
}

func GetTemplateByID(templateID string) (map[string]interface{}, error) {
	return dataTemplateRepository.GetTemplateByID(templateID)
}

func GetTemplatesByPlotID(plotID string) ([]map[string]interface{}, error) {
	return dataTemplateRepository.GetTemplatesByPlotID(plotID)
}

// func GetHTMLByPlotID(templateID string) (model map[string]interface{}, error) {
// 	return workflowRepository.GetWorkflowsByID(workflowsID)
// }

func GetLastTemplatesByPlotID(plotID string) (map[string]interface{}, error) {
	return dataTemplateRepository.GetLastTemplateByPlotID(plotID)
}

func GetTemplatesByUser(userID string) ([]map[string]interface{}, error) {
	return dataTemplateRepository.GetTemplateByUser(userID)
}

func GetTemplatePagination(paginationData requestDtos.TemplateForMatrixView) (model.DataPaginatedresponse, error) {
	filter := bson.M{
		"userid": paginationData.UserID,
	}
	projection := GetDynamicProjectionForTemplate(paginationData.Fields)

	var data []map[string]interface{}

	response, err := dataTemplateRepository.GetTemplatePaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "dataTemplate", "userid", data, paginationData.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred:", err.Error())
		return model.DataPaginatedresponse{}, err
	}

	return model.DataPaginatedresponse(response), nil
}

func GetDynamicProjectionForTemplate(fields []string) bson.D {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "userid", Value: 1},
	}

	for _, field := range fields {
		projection = append(projection, bson.E{Key: field, Value: 1})
	}
	return projection
}
