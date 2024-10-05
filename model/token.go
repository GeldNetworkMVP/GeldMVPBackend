package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tokens struct {
	TokenID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	PlotID      string             `json:"plotid" bson:"plotid"`
	TokenName   string             `json:"tokenname" bson:"tokenname"`
	Description string             `json:"description" bson:"description"`
	CID         string             `json:"cid" bson:"cid"`
	Price       string             `json:"price" bson:"price"`
	IPFSStatus  string             `json:"status" bson:"status"`
	BCStatus    string             `json:"bcstatus" bson:"bcstatus"`
	TokenHash   string             `json:"tokenhash" bson:"tokenhash"`
}

type TokenPayload struct {
	TokenID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	PlotID      string             `json:"plotid" bson:"plotid"`
	TokenName   string             `json:"tokenname" bson:"tokenname"`
	Description string             `json:"description" bson:"description"`
	Price       string             `json:"price" bson:"price"`
	FileType    string             `json:"filetype" bson:"filetype"`
	BCStatus    string             `json:"bcstatus" bson:"bcstatus"`
	BCHash      string             `json:"bchash" bson:"bchash"`
}

type TokenPaginatedresponse struct {
	Content        []Tokens `json:"content" bson:"content" validate:"required"`
	PaginationInfo PaginationTemplate
}

type TokenTransactions struct {
	TransactionID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	TransactionStatus string             `json:"status" bson:"status"`
	TXNHash           string             `json:"txnhash" bson:"txnhash"`
	PlotID            string             `json:"plotid" bson:"plotid"`
	TokenID           string             `json:"tokenid" bson:"tokenid"`
	DBStatus          string             `json:"dbstatus" bson:"dbstatus"`
}
