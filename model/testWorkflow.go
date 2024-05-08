package model

type PlotChainData struct {
	PlotID   string
	Workflow Workflow
	Status   string
}

type Workflow struct {
	WorkflowID string
	Status     string
	Stages     Stages
}

type Stages struct {
	StageID         string
	StageName       string
	StageFieldNames StageFieldNames
}

type StageFieldNames struct {
	Field1 string
	Field2 string
	Field3 string
}
