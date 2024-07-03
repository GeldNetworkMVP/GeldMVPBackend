package dataConfigRepository

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories"
)

type DataTemplateRepository struct{}

var DataTemplate = "dataTemplate"

func (r *DataTemplateRepository) SaveDataTemplate(model map[string]interface{}) (string, error) {
	return repositories.SaveDynamicData(model, DataTemplate)
}
