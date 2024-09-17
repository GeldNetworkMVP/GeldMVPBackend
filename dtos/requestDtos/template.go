package requestDtos

type TemplateForMatrixView struct {
	UserID        string   `json:"userid" bson:"userid"`
	Fields        []string `json:"fields" bson:"fields"`
	PageSize      int32    `json:"pagesize" bson:"pagesize"`
	RequestedPage int32    `json:"requestedPage" bson:"requestedPage" `
	SortbyField   string   `json:"sortbyfield" bson:"sortbyfield" `
	SortType      int      `json:"sorttype" bson:"sorttype"`
}
