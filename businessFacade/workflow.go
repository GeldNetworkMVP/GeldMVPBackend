package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

func CreateWorkflows(workflows model.Workflows) (string, error) {
	return workflowRepository.CreateWorkflow(workflows)
}
