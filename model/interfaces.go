package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveType interface {
	Workflows | MasterData | Stages | DataCollection | AppUser
}

type FindOneType interface {
	string | primitive.ObjectID
}
