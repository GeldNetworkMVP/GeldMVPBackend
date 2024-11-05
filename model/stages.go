package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stages struct {
	StageID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	// UserID      string             `json:"userid" bson:"userid"`
	StageName   string   `json:"stagename" bson:"stagename"`
	Description string   `json:"description" bson:"description"`
	Fields      []Fields `json:"fields" bson:"fields"`
}

type Fields struct {
	ValueKey  string `json:"valuekey" bson:"valuekey"`
	ValueType string `json:"valuetype" bson:"valuetype"`
}

type StagePaginatedresponse struct {
	Content        []Stages `json:"content" bson:"content" validate:"required"`
	PaginationInfo PaginationTemplate
}

type StagesNames struct {
	StageArray []string `json:"stagearray" bson:"stagearray"`
}
