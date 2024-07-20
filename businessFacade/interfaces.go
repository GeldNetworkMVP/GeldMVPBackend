package businessFacade

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories/dataConfigRepository"
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories/tokenGeldRepository"
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories/userAndPermissions"
)

var (
	workflowRepository     dataConfigRepository.WorkflowRepository
	stageRepository        dataConfigRepository.StageRepository
	masterdataRepository   dataConfigRepository.MasterDataRepository
	dataTemplateRepository dataConfigRepository.DataTemplateRepository
	userRepository         userAndPermissions.UserRepository
	tokensRepository       tokenGeldRepository.TokenRepository
)
