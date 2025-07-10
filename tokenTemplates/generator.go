package tokenTemplates

import (
	"bytes"
	"fmt"
	"text/template"
	// "github.com/GeldNetworkMVP/GeldMVPBackend/services"
)

type Stage struct {
	ID          string `json:"_id"`
	Description string `json:"description"`
	StageName   string `json:"stagename"`
}

type Plot struct {
	ID             string   `json:"_id"`
	CollectionName string   `json:"collectionname"`
	Contact        string   `json:"contact"`
	DataID         string   `json:"dataid"`
	Description    string   `json:"description"`
	Location       string   `json:"location"`
	Purpose        []string `json:"purpose"`
	Type           string   `json:"type"`
	UserID         string   `json:"userid"`
}

type Officer struct {
	ID             string   `json:"_id"`
	CollectionName string   `json:"collectionname"`
	Contact        string   `json:"contact"`
	DataID         string   `json:"dataid"`
	Description    string   `json:"description"`
	Location       string   `json:"location"`
	Purpose        []string `json:"purpose"`
	Type           string   `json:"type"`
	UserID         string   `json:"userid"`
}

type CollectionData struct {
	CollectionName string
	Key            string
	Data           interface{}
}

type TemplateData struct {
	TemplateName string
	Key          string
	Data         interface{}
}

type PageData struct {
	Collections map[string]CollectionData
	Templates   map[string]TemplateData
}

// var (
// 	documentStart = services.ReadFromFile("templates/svgHeader.txt")
// 	documentEnd   = services.ReadFromFile("templates/svgFooter.txt")
// 	style         = services.ReadFromFile("templates/style.scss")
// 	styleStart    = `<style>`
// 	styleEnd      = `</style>`
// 	svg           = services.ReadFromFile("services/svgGeneratorService/templates/temp.svg")
// )

func GenerateTokenPreview(htmlData []map[string]interface{}) (string, error) {
	uniqueCollections := extractUniqueCollections(htmlData)

	uniqueTemplates := extractUniqueTemplates(htmlData)

	html, err := GenerateSVG(uniqueCollections, uniqueTemplates)
	if err != nil {
		fmt.Printf("Error creating HTML")
	}
	return html, nil
}

func extractUniqueCollections(data []map[string]interface{}) map[string]CollectionData {
	uniqueCollections := make(map[string]CollectionData)

	for _, item := range data {
		extractCollections(item, uniqueCollections)
	}

	return uniqueCollections
}

func extractCollections(item map[string]interface{}, uniqueCollections map[string]CollectionData) {
	for key, value := range item {
		if nestedMap, ok := value.(map[string]interface{}); ok {
			if collectionName, exists := nestedMap["collectionname"]; exists {
				name := collectionName.(string)
				if _, alreadyExists := uniqueCollections[name]; !alreadyExists {
					uniqueCollections[name] = CollectionData{
						CollectionName: name,
						Key:            key,
						Data:           nestedMap,
					}
				}
			}
			extractCollections(nestedMap, uniqueCollections)
		}
	}
}

func extractUniqueTemplates(data []map[string]interface{}) map[string]TemplateData {
	uniqueTemplates := make(map[string]TemplateData)

	for _, item := range data {
		extractTemplates(item, uniqueTemplates)
	}

	return uniqueTemplates
}

func extractTemplates(item map[string]interface{}, uniqueTemplates map[string]TemplateData) {
	for key, value := range item {
		if nestedMap, ok := value.(map[string]interface{}); ok {
			if templateName, exists := nestedMap["templatename"]; exists {
				name := templateName.(string)
				if _, alreadyExists := uniqueTemplates[name]; !alreadyExists {
					uniqueTemplates[name] = TemplateData{
						TemplateName: name,
						Key:          key,
						Data:         nestedMap,
					}
				}
			}
			extractTemplates(nestedMap, uniqueTemplates)
		} else if templateName, exists := item["templatename"]; exists {
			name := templateName.(string)
			if _, alreadyExists := uniqueTemplates[name]; !alreadyExists {
				uniqueTemplates[name] = TemplateData{
					TemplateName: name,
					Key:          key,
					Data:         item,
				}
			}
		}
	}
}

func GenerateSVG(collections map[string]CollectionData, templates map[string]TemplateData) (string, error) {
	// 1. Parse templates
	tmpl := template.Must(template.ParseFiles(
		"templates/base.tmpl",
		"templates/header.tmpl",
		"templates/footer.tmpl",
		"templates/token.tmpl",
	))

	// 2. Prepare data
	data := PageData{
		Collections: collections,
		Templates:   templates,
	}

	// 3. Execute template directly
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "base", data)
	if err != nil {
		return "", err
	}

	// 4. Debug prints (will work now)
	fmt.Println("the HTML ", buf.String())

	return buf.String(), nil
}
