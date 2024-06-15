package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories/dataConfigRepository"
)

var (
	workflowRepository   dataConfigRepository.WorkflowRepository
	stageRepository      dataConfigRepository.StageRepository
	masterdataRepository dataConfigRepository.MasterDataRepository
)
