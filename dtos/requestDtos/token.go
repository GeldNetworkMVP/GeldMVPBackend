package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateToken struct {
	TokenID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Status  string             `json:"status" bson:"status"`
}

type TokenForMatrixView struct {
	Status        string `json:"status" bson:"status"`
	PageSize      int32  `json:"pagesize" bson:"pagesize"`
	RequestedPage int32  `json:"requestedPage" bson:"requestedPage" `
	SortbyField   string `json:"sortbyfield" bson:"sortbyfield" `
	SortType      int    `json:"sorttype" bson:"sorttype"`
}
