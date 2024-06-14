package dataConfigRepository

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

type WorkflowRepository struct{}

var Workflow = "workflows"

func (r *WorkflowRepository) CreateWorkflow(workflow model.Workflows) (string, error) {
	return repositories.Save(workflow, Workflow)
}
