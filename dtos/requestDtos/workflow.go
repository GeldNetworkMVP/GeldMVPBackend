package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateWorkflow struct {
	WorkflowID   primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	WorkflowName string             `json:"workflowname" bson:"workflowname"`
	Description  string             `json:"description" bson:"description"`
	Stages       []string           `json:"stages" bson:"stages"`
}