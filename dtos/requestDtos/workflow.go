package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateWorkflow struct {
	WorkflowID   primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID       string             `json:"userid" bson:"userid"`
	WorkflowName string             `json:"workflowname" bson:"workflowname"`
	Description  string             `json:"description" bson:"description"`
	Stages       []string           `json:"stages" bson:"stages"`
}

type WorkflowForMatrixView struct {
	UserID        string `json:"userid" bson:"userid"`
	PageSize      int32  `json:"pagesize" bson:"pagesize"`
	RequestedPage int32  `json:"requestedPage" bson:"requestedPage" `
	SortbyField   string `json:"sortbyfield" bson:"sortbyfield" `
	SortType      int    `json:"sorttype" bson:"sorttype"`
}
