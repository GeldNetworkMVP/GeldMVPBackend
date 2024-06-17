package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveType interface {
	Workflows | MasterData | Stages | DataCollection
}

type FindOneType interface {
	string | primitive.ObjectID
}
