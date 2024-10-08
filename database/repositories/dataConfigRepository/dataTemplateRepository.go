package dataConfigRepository

import (
	"context"
	"errors"

	"github.com/GeldNetworkMVP/GeldMVPBackend/database/connections"
	"github.com/GeldNetworkMVP/GeldMVPBackend/database/repositories"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataTemplateRepository struct{}

var DataTemplate = "dataTemplate"

func (r *DataTemplateRepository) SaveDataTemplate(model map[string]interface{}) (string, error) {
	return repositories.SaveDynamicData(model, DataTemplate, "templatename")
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

	// objectID, err := primitive.ObjectIDFromHex(plotID)
	// if err != nil {
	// 	logs.ErrorLogger.Println("Invalid plotID:", err.Error())
	// 	return nil, err
	// }

	filter := bson.M{"plot._id": plotID}

	// Execute the query
	cursor, err := connections.GetSessionClient(DataTemplate).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

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

func (r *DataTemplateRepository) GetLastTemplateByPlotID(plotID string) (map[string]interface{}, error) {
	ctx := context.TODO()

	opts := options.FindOne().SetSort(bson.D{{Key: "_id", Value: -1}})
	var result map[string]interface{}
	err := connections.GetSessionClient(DataTemplate).FindOne(ctx, bson.M{"plotid": plotID}, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logs.ErrorLogger.Println("No template found for plotID:", plotID)
			return nil, nil
		}
		logs.ErrorLogger.Println("Error retrieving template:", err.Error())
		return nil, err
	}

	return result, nil
}

func (r *DataTemplateRepository) GetTemplateByUser(userID string) ([]map[string]interface{}, error) {
	ctx := context.TODO()
	cursor, err := connections.GetSessionClient(DataTemplate).Find(ctx, bson.M{"userid": userID})
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

func (r *DataTemplateRepository) GetTemplatePaginatedResponse(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, data []map[string]interface{}, sort int) (model.DataPaginatedresponse, error) {
	contentResponse, paginationResponse, err := repositories.PaginateResponse[[]map[string]interface{}](
		filterConfig,
		projectionData,
		pagesize,
		pageNo,
		collectionName,
		sortingFeildName,
		data,
		sort,
	)
	var response model.DataPaginatedresponse
	if err != nil {
		return response, err
	}
	response.Content = contentResponse
	response.PaginationInfo = paginationResponse
	return response, nil
}
