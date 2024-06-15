package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateStages struct {
	StageID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	StageName   string             `json:"stagename" bson:"stagename"`
	Description string             `json:"description" bson:"description"`
	Fields      []string           `json:"fields" bson:"fields"`
}
