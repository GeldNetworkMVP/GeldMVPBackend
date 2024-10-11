package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MasterData struct {
	DataID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	// UserID      string             `json:"userid" bson:"userid"`
	DataName    string `json:"dataname" bson:"dataname"`
	Description string `json:"description" bson:"description"`
	//DataCollection []DataCollection   `json:"dataCollection" bson:"dataCollection"`
	MasterDataFields   []string `json:"mfields" bson:"mfields"`
	RecordsInContainer string   `json:"noOfRecords" bson:"noOfRecords"`
}

type DataCollection struct {
	CollectionID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	DataID       string             `json:"dataid" bson:"dataid"`
	// UserID         string             `json:"userid" bson:"userid"`
	CollectionName string   `json:"collectionname" bson:"collectionname"`
	Description    string   `json:"description" bson:"description"`
	Purpose        []string `json:"purpose" bson:"purpose"`
	Location       string   `json:"location" bson:"location"`
	Contact        string   `json:"contact" bson:"contact"`
	Type           string   `json:"type" bson:"type"`
}

type MDataPaginatedresponse struct {
	Content        []MasterData `json:"content" bson:"content" validate:"required"`
	PaginationInfo PaginationTemplate
}

type DataPaginatedresponse struct {
	Content        []map[string]interface{} `json:"content" bson:"content" validate:"required"`
	PaginationInfo PaginationTemplate
}
