package dataConfigRepository

import (
	"context"

	"github.com/GeldNetworkMVP/GeldMVPBackend/database/connections"
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories"
	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/requestDtos"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WorkflowRepository struct{}

var Workflow = "workflows"

func (r *WorkflowRepository) CreateWorkflow(workflow model.Workflows) (string, error) {
	return repositories.Save(workflow, Workflow)
}

func (r *WorkflowRepository) GetWorkflowsByID(workflowID string) (model.Workflows, error) {
	var workflow model.Workflows
	objectId, err := primitive.ObjectIDFromHex(workflowID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetWorkflowsByID : workflowRepository: ", err.Error())
	}
	rst, err := connections.GetSessionClient("workflows").Find(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return workflow, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&workflow)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection document in GetWorkflowByID:workflowRepository.go: ", err.Error())
			return workflow, err
		}
	}
	return workflow, err
}

func (r *WorkflowRepository) UpdateWorkflows(UpdateObject requestDtos.UpdateWorkflow, update primitive.M) (model.Workflows, error) {
	var WorkflowUpdateResponse model.Workflows
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("workflows").FindOneAndUpdate(context.TODO(), bson.M{"_id": UpdateObject.WorkflowID}, update, &opt)
	if rst != nil {
		err := rst.Decode((&WorkflowUpdateResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error Occured while Update Workflow", err.Error())
			return WorkflowUpdateResponse, err
		}
		return WorkflowUpdateResponse, err
	}
	return WorkflowUpdateResponse, nil
}
