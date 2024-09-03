package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Workflows struct {
	WorkflowID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	// UserID       string             `json:"userid" bson:"userid"`
	WorkflowName string   `json:"workflowname" bson:"workflowname"`
	Description  string   `json:"description" bson:"description"`
	Stages       []string `json:"stages" bson:"stages"`
}

type WorkflowPaginatedresponse struct {
	Content        []Workflows `json:"content" bson:"content" validate:"required"`
	PaginationInfo PaginationTemplate
}
