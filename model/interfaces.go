package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveType interface {
	Workflows | MasterData | Stages | DataCollection | AppUser | Tokens | TokenTransactions
}

type FindOneType interface {
	string | primitive.ObjectID
}
