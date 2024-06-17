package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateWorkflows(workflows model.Workflows) (string, error) {
	return workflowRepository.CreateWorkflow(workflows)
}

func GetWorkflowsByID(workflowsID string) (model.Workflows, error) {
	return workflowRepository.GetWorkflowsByID(workflowsID)
}

func UpdateWorkflow(UpdateObject requestDtos.UpdateWorkflow) (model.Workflows, error) {
	update := bson.M{
		"$set": bson.M{"workflowname": UpdateObject.WorkflowName, "description": UpdateObject.Description, "stages": UpdateObject.Stages},
	}
	return workflowRepository.UpdateWorkflows(UpdateObject, update)
}

func DeleteWorkflowByID(WorkflowID string) error {
	return workflowRepository.DeleteWorkflow(WorkflowID)
}

func GetWorkflowDataPagination(paginationData requestDtos.WorkflowForMatrixView) (model.WorkflowPaginatedresponse, error) {
	filter := bson.M{
		"userid": paginationData.UserID,
	}
	projection := GetProjectionDataMatrixViewForWorkflowData()
	var data []model.Workflows
	response, err := workflowRepository.GetWorkflowDataPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "workflows", "userid", data, paginationData.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return model.WorkflowPaginatedresponse(response), err
	}
	return model.WorkflowPaginatedresponse(response), err
}

func GetProjectionDataMatrixViewForWorkflowData() bson.D {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "userid", Value: 1},
		{Key: "workflowname", Value: 1},
		{Key: "description", Value: 1},
		{Key: "stages", Value: 1},
	}
	return projection
}
