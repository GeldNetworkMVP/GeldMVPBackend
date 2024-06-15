package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
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
