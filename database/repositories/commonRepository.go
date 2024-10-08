package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/GeldNetworkMVP/GeldMVPBackend/commons"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbName = commons.GoDotEnvVariable("DB_NAME")

func Save[T model.SaveType](model T, collection string) (string, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(commons.GoDotEnvVariable("DB_URI")))
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	}
	defer client.Disconnect(context.Background())

	session, err := client.StartSession()
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	}
	defer session.EndSession(context.Background())

	rst, err := session.Client().Database(DbName).Collection(collection).InsertOne(context.TODO(), model)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	}
	id := rst.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func SaveDynamicData(model map[string]interface{}, collection string, check string) (string, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(commons.GoDotEnvVariable("DB_URI")))
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	}
	defer client.Disconnect(context.Background())

	session, err := client.StartSession()
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	}
	defer session.EndSession(context.Background())

	data, ok := model[check].(string)
	if !ok {
		return "", fmt.Errorf("name field is missing or not a string")
	}

	filter := bson.M{check: data}
	var existingDoc bson.M
	collectionRef := session.Client().Database(DbName).Collection(collection)
	err = collectionRef.FindOne(context.TODO(), filter).Decode(&existingDoc)
	if err != nil && err != mongo.ErrNoDocuments {
		logs.ErrorLogger.Println("Error while checking existing template:", err.Error())
		return "", err
	}

	if existingDoc != nil {
		return "", fmt.Errorf("a template with the name '%s' already exists", data)
	}

	rst, err := collectionRef.InsertOne(context.TODO(), model)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	}

	id := rst.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil

}

func FindById(idName string, id string, collection string) (*mongo.Cursor, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(commons.GoDotEnvVariable("DB_URI")))
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil, err
	}
	defer client.Disconnect(context.Background())

	session, err := client.StartSession()
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil, err
	}
	defer session.EndSession(context.Background())

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	findOptions.SetProjection(bson.M{"otp": 0})
	rst, err := session.Client().Database(DbName).Collection(collection).Find(context.TODO(), bson.D{{idName, id}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindOne[T model.FindOneType](idName string, id T, collection string) *mongo.SingleResult {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(commons.GoDotEnvVariable("DB_URI")))
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil
	}
	defer client.Disconnect(context.Background())

	session, err := client.StartSession()
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil
	}
	defer session.EndSession(context.Background())
	rst := session.Client().Database(DbName).Collection(collection).FindOne(context.TODO(), bson.D{{idName, id}})
	return rst
}

func FindById1AndId2(idName1 string, id1 string, idName2 string, id2 string, collection string) (*mongo.Cursor, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(commons.GoDotEnvVariable("DB_URI")))
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil, err
	}
	defer client.Disconnect(context.Background())

	session, err := client.StartSession()
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil, err
	}
	defer session.EndSession(context.Background())
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := session.Client().Database(DbName).Collection(collection).Find(context.TODO(), bson.D{{idName1, id1}, {idName2, id2}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindById1Id2Id3(idName1 string, id1 string, idName2 string, id2 string, idName3 string, id3 string, collection string) (*mongo.Cursor, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(commons.GoDotEnvVariable("DB_URI")))
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil, err
	}
	defer client.Disconnect(context.Background())

	session, err := client.StartSession()
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil, err
	}
	defer session.EndSession(context.Background())
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := session.Client().Database(DbName).Collection(collection).Find(context.TODO(), bson.D{{idName1, id1}, {idName2, id2}, {idName3, id3}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindByFieldInMultipleValus(fields string, tags []string, collection string) (*mongo.Cursor, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(commons.GoDotEnvVariable("DB_URI")))
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil, err
	}
	defer client.Disconnect(context.Background())

	session, err := client.StartSession()
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil, err
	}
	defer session.EndSession(context.Background())
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := session.Client().Database(DbName).Collection(collection).Find(context.TODO(), bson.D{{fields, bson.D{{"$in", tags}}}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindOneAndUpdate(findBy string, value string, update primitive.M, projectionData primitive.M, collection string) *mongo.SingleResult {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(commons.GoDotEnvVariable("DB_URI")))
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil
	}
	defer client.Disconnect(context.Background())

	session, err := client.StartSession()
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nil
	}
	defer session.EndSession(context.Background())
	after := options.After
	projection := projectionData
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Projection:     &projection,
	}
	rst := session.Client().Database(DbName).Collection(collection).FindOneAndUpdate(context.TODO(), bson.M{findBy: value}, update, &opt)
	return rst
}

func Remove(idName string, id, collection string) (int64, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(commons.GoDotEnvVariable("DB_URI")))
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return 0, err
	}
	defer client.Disconnect(context.Background())

	session, err := client.StartSession()
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return 0, err
	}
	defer session.EndSession(context.Background())
	result, err := session.Client().Database(DbName).Collection(collection).DeleteMany(context.TODO(), bson.M{idName: id})
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return 0, err
	}
	return result.DeletedCount, nil
}

type paginateResponseType interface {
	[]model.Workflows | []model.MasterData | []model.Stages | []model.DataCollection | []model.AppUser | []model.Tokens | []map[string]interface{}
}

func PaginateResponse[PaginatedData paginateResponseType](filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, object PaginatedData, sort int) (PaginatedData, model.PaginationTemplate, error) {
	var paginationdata model.PaginationTemplate
	ctx := context.Background()
	connectionString := os.Getenv("DB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		logs.ErrorLogger.Println("failed to connect to DB: ", err.Error())
	}
	defer client.Disconnect(ctx)
	dbConnection := client.Database(DbName)
	filter := filterConfig
	limit := int64(pagesize)
	page := int64(pageNo)
	collection := dbConnection.Collection(collectionName)
	projection := projectionData
	paginatedData, paginateerr := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Sort(sortingFeildName, sort).Select(projection).Filter(filter).Decode(&object).Find()
	paginationdata.TotalElements = int32(paginatedData.Pagination.Total)
	paginationdata.TotalPages = int32(paginatedData.Pagination.TotalPage)
	paginationdata.Currentpage = int32(paginatedData.Pagination.Page)
	paginationdata.PageSize = int32(paginatedData.Pagination.PerPage)
	paginationdata.Previouspage = int32(paginatedData.Pagination.Prev)
	paginationdata.NextPage = int32(paginatedData.Pagination.Next)
	if paginateerr != nil {
		logs.ErrorLogger.Println("Pagination failure :", paginateerr.Error())
		return object, paginationdata, err
	}
	return object, paginationdata, nil
}

// func PaginateWithCustomSort[PaginatedData paginateResponseType](filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, sortyType int, object PaginatedData) (PaginatedData, model.PaginationTemplate, error) {
// 	var paginationdata model.PaginationTemplate
// 	ctx := context.Background()
// 	connectionString := os.Getenv("DB_URI")
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
// 	if err != nil {
// 		logs.ErrorLogger.Println("failed to connect to DB: ", err.Error())
// 	}
// 	defer client.Disconnect(ctx)
// 	dbConnection := client.Database(DbName)
// 	filter := filterConfig
// 	limit := int64(pagesize)
// 	page := int64(pageNo)
// 	collection := dbConnection.Collection(collectionName)
// 	projection := projectionData
// 	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Sort(sortingFeildName, sortyType).Select(projection).Filter(filter).Decode(&object).Find()
// 	paginationdata.TotalElements = int32(paginatedData.Pagination.Total)
// 	paginationdata.TotalPages = int32(paginatedData.Pagination.TotalPage)
// 	paginationdata.Currentpage = int32(paginatedData.Pagination.Page)
// 	paginationdata.PageSize = int32(paginatedData.Pagination.PerPage)
// 	paginationdata.Previouspage = int32(paginatedData.Pagination.Prev)
// 	paginationdata.NextPage = int32(paginatedData.Pagination.Next)
// 	if err != nil {
// 		logs.ErrorLogger.Println("Pagination failure :", err.Error())
// 		return object, paginationdata, err
// 	}
// 	return object, paginationdata, nil
// }

func TestPaginateResponse[PaginatedData paginateResponseType](filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, object PaginatedData, sort int) (PaginatedData, model.PaginationTemplate, error) {
	var paginationdata model.PaginationTemplate
	ctx := context.Background()
	connectionString := os.Getenv("DB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		logs.ErrorLogger.Println("failed to connect to DB: ", err.Error())
	}
	defer client.Disconnect(ctx)
	dbConnection := client.Database(DbName)
	filter := filterConfig
	limit := int64(pagesize)
	page := int64(pageNo)
	collection := dbConnection.Collection(collectionName)
	projection := projectionData

	paginator := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Select(projection).Filter(filter)

	if sortingFeildName != "" {
		paginator = paginator.Sort(sortingFeildName, sort)
	}

	paginatedData, paginateerr := paginator.Decode(&object).Find()

	if paginateerr != nil {
		logs.ErrorLogger.Println("Pagination failure:", paginateerr.Error())
		return object, paginationdata, paginateerr
	}

	paginationdata.TotalElements = int32(paginatedData.Pagination.Total)
	paginationdata.TotalPages = int32(paginatedData.Pagination.TotalPage)
	paginationdata.Currentpage = int32(paginatedData.Pagination.Page)
	paginationdata.PageSize = int32(paginatedData.Pagination.PerPage)
	paginationdata.Previouspage = int32(paginatedData.Pagination.Prev)
	paginationdata.NextPage = int32(paginatedData.Pagination.Next)
	if paginateerr != nil {
		logs.ErrorLogger.Println("Pagination failure :", paginateerr.Error())
		return object, paginationdata, err
	}
	return object, paginationdata, nil
}
