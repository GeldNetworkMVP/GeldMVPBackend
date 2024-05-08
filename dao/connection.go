package dao

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/commons"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get db name from .env file
var DbName = commons.GoDotEnvVariable("DB_NAME")

/*
Connection The Mgo Connection
@author - Azeem Ashraf
*/
type Connection struct {
}

func (cd *Connection) Connect() (mongo.Session, error) {
	return commons.GetMongoSession()
}
