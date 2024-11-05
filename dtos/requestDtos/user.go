package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateUser struct {
	AppUserID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	// AdminUserID string             `json:"userid" bson:"userid"`
	Email       string `json:"email" bson:"email"`
	Contact     string `json:"contact" bson:"contact"`
	Designation string `json:"designation" bson:"designation"`
	// EncPW       string `json:"encpw" bson:"encpw"`
	Status string `json:"status" bson:"status"`
}
type UpdateUserStatus struct {
	AppUserID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Status    string             `json:"status" bson:"status"`
}

type UserForMatrixView struct {
	// UserID        string `json:"userid" bson:"userid"`
	PageSize      int32  `json:"pagesize" bson:"pagesize"`
	RequestedPage int32  `json:"requestedPage" bson:"requestedPage" `
	SortbyField   string `json:"sortbyfield" bson:"sortbyfield" `
	SortType      int    `json:"sorttype" bson:"sorttype"`
}

type UpdateUserPublicKey struct {
	Email     string `json:"email" bson:"email"`
	PublicKey string `json:"publickey" bson:"publickey"`
}
