package businessFacade

func SaveDataTemplate(model map[string]interface{}) (string, error) {
	return dataTemplateRepository.SaveDataTemplate(model)
}
