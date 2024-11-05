package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppUser struct {
	AppUserID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	// AdminUserID string             `json:"userid" bson:"userid"`
	Email       string `json:"email" bson:"email"`
	Contact     string `json:"contact" bson:"contact"`
	Designation string `json:"designation" bson:"designation"`
	EncPW       []byte `json:"encpw" bson:"encpw"`
	Status      string `json:"status" bson:"status"`
	Company     string `json:"company" bson:"company"`
	Username    string `json:"username" bson:"username"`
	PublicKey   string `json:"publickey" bson:"publickey"`
}

type UserPayload struct {
	AppUserID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	// AdminUserID string             `json:"userid" bson:"userid"`
	Username    string `json:"username" bson:"username"`
	Email       string `json:"email" bson:"email"`
	Contact     string `json:"contact" bson:"contact"`
	Designation string `json:"designation" bson:"designation"`
	Password    string `json:"encpw" bson:"encpw"`
	Status      string `json:"status" bson:"status"`
	Company     string `json:"company" bson:"company"`
}

type UserPaginatedResponse struct {
	Content        []AppUserDetails `json:"content" bson:"content" validate:"required"`
	PaginationInfo PaginationTemplate
}

type AppUserDetails struct {
	AppUserID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	// AdminUserID string             `json:"userid" bson:"userid"`
	Email       string `json:"email" bson:"email"`
	Contact     string `json:"contact" bson:"contact"`
	Designation string `json:"designation" bson:"designation"`
	Status      string `json:"status" bson:"status"`
	Company     string `json:"company" bson:"company"`
	PublicKey   string `json:"publickey" bson:"publickey"`
	Username    string `json:"username" bson:"username"`
}

type AppCredentials struct {
	Email string `json:"email" bson:"email"`
	Pw    string `json:"pw" bson:"pw"`
}

type UserExistence struct {
	Status    string `json:"status" bson:"status"`
	Operative string `json:"op" bson:"op"`
}
