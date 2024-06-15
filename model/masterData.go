package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MasterData struct {
	DataID      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	DataName    string             `json:"dataname" bson:"dataname"`
	Description string             `json:"description" bson:"description"`
	//DataCollection []DataCollection   `json:"dataCollection" bson:"dataCollection"`
}

type DataCollection struct {
	CollectionID   primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	DataID         string             `json:"dataid" bson:"dataid"`
	CollectionName string             `json:"collectionname" bson:"collectionname"`
	Description    string             `json:"description" bson:"description"`
	Purpose        []string           `json:"purpose" bson:"purpose"`
	Location       string             `json:"location" bson:"location"`
	Contact        string             `json:"contact" bson:"contact"`
	Type           string             `json:"type" bson:"type"`
}
