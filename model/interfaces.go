package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveType interface {
	Workflows
}

type FindOneType interface {
	string | primitive.ObjectID
}
