package model

type ManageData struct {
	Name           string
	Value          string
	Source_account string
	Asset_code     string
	Amount         string
	To             string
	From           string
}

type Keys struct {
	PK string
	SK []byte
}

type IssuerResponse struct {
	IssuerPK string
	Result   string
}
