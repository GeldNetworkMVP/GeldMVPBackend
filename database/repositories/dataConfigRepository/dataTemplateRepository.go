package dataConfigRepository

import (
	"context"
	"errors"

	"github.com/GeldNetworkMVP/GeldMVPBackend/database/connections"
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataTemplateRepository struct{}

var DataTemplate = "dataTemplate"

func (r *DataTemplateRepository) SaveDataTemplate(model map[string]interface{}) (string, error) {
	return repositories.SaveDynamicData(model, DataTemplate)
}

func (r *DataTemplateRepository) GetTemplateByID(templateID string) (map[string]interface{}, error) {
	objectId, err := primitive.ObjectIDFromHex(templateID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetTemplateByID : datatemplateRepository: ", err.Error())
	}
	ctx := context.TODO()
	result := bson.M{}
	err = connections.GetSessionClient(DataTemplate).FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Template not found")
		}
		logs.ErrorLogger.Println("Error retrieving template:", err.Error())
		return nil, err
	}
	return result, nil
}

func (r *DataTemplateRepository) GetTemplatesByPlotID(plotID string) ([]map[string]interface{}, error) {
	ctx := context.TODO()
	cursor, err := connections.GetSessionClient(DataTemplate).Find(ctx, bson.M{"plotid": plotID})
	if err != nil {
		return nil, err
	}

	var templates []map[string]interface{}

	for cursor.Next(ctx) {
		var result map[string]interface{}
		err := cursor.Decode(&result)
		if err != nil {
			logs.ErrorLogger.Println("Error retrieving template:", err.Error())
			return nil, err
		}
		templates = append(templates, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return templates, nil
}
