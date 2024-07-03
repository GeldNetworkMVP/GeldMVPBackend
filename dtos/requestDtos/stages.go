package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateStages struct {
	StageID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID      string             `json:"userid" bson:"userid"`
	StageName   string             `json:"stagename" bson:"stagename"`
	Description string             `json:"description" bson:"description"`
	Fields      []string           `json:"fields" bson:"fields"`
}

type StagesForMatrixView struct {
	UserID        string `json:"userid" bson:"userid"`
	PageSize      int32  `json:"pagesize" bson:"pagesize"`
	RequestedPage int32  `json:"requestedPage" bson:"requestedPage" `
	SortbyField   string `json:"sortbyfield" bson:"sortbyfield" `
	SortType      int    `json:"sorttype" bson:"sorttype"`
}
