package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stages struct {
	StageID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID      string             `json:"userid" bson:"userid"`
	StageName   string             `json:"stagename" bson:"stagename"`
	Description string             `json:"description" bson:"description"`
	Fields      []string           `json:"fields" bson:"fields"`
}

type StagePaginatedresponse struct {
	Content        []Stages `json:"content" bson:"content" validate:"required"`
	PaginationInfo PaginationTemplate
}
