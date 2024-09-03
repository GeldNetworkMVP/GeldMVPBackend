package businessFacade

func SaveDataTemplate(model map[string]interface{}) (string, error) {
	return dataTemplateRepository.SaveDataTemplate(model)
}

func GetTemplateByID(templateID string) (map[string]interface{}, error) {
	return dataTemplateRepository.GetTemplateByID(templateID)
}

func GetTemplatesByPlotID(plotID string) ([]map[string]interface{}, error) {
	return dataTemplateRepository.GetTemplatesByPlotID(plotID)
}

// func GetHTMLByPlotID(templateID string) (model map[string]interface{}, error) {
// 	return workflowRepository.GetWorkflowsByID(workflowsID)
// }

func GetLastTemplatesByPlotID(plotID string) (map[string]interface{}, error) {
	return dataTemplateRepository.GetLastTemplateByPlotID(plotID)
}
