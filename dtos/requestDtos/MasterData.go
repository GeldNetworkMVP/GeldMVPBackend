package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateMasterData struct {
	DataID      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID      string             `json:"userid" bson:"userid"`
	DataName    string             `json:"dataname" bson:"dataname"`
	Description string             `json:"description" bson:"description"`
	//DataCollection UpdateDataCollection `json:"dataCollection" bson:"dataCollection"`
}

type UpdateDataCollection struct {
	CollectionID   primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	DataID         string             `json:"dataid" bson:"dataid"`
	UserID         string             `json:"userid" bson:"userid"`
	CollectionName string             `json:"collectionname" bson:"collectionname"`
	Description    string             `json:"description" bson:"description"`
	Purpose        []string           `json:"purpose" bson:"purpose"`
	Location       string             `json:"location" bson:"location"`
	Contact        string             `json:"contact" bson:"contact"`
	Type           string             `json:"type" bson:"type"`
}

type MasterDataForMatrixView struct {
	UserID        string `json:"userid" bson:"userid"`
	PageSize      int32  `json:"pagesize" bson:"pagesize"`
	RequestedPage int32  `json:"requestedPage" bson:"requestedPage" `
	SortbyField   string `json:"sortbyfield" bson:"sortbyfield" `
	SortType      int    `json:"sorttype" bson:"sorttype"`
}

type DataRecordForMatrixView struct {
	DataID        string `json:"dataid" bson:"dataid"`
	PageSize      int32  `json:"pagesize" bson:"pagesize"`
	RequestedPage int32  `json:"requestedPage" bson:"requestedPage" `
	SortbyField   string `json:"sortbyfield" bson:"sortbyfield" `
	SortType      int    `json:"sorttype" bson:"sorttype"`
}
