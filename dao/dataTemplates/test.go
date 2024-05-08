package datatemplates

import (
	"context"
	"log"

	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

func (cd *Connection) SaveWorkflow(stellarNFT model.NFTWithTransaction, marketPlaceNFT model.MarketPlaceNFT) (error, error) {
	session, err := cd.connect()
	if err != nil {
		log.Println("Error when connecting to DB " + err.Error())
	}
	defer session.EndSession(context.TODO())
	c := session.Client().Database(dbName).Collection("NFTStellar")
	c2 := session.Client().Database(dbName).Collection("MarketPlaceNFT")
	_, err2 := c.InsertOne(context.TODO(), stellarNFT)
	_, err = c2.InsertOne(context.TODO(), marketPlaceNFT)
	if err != nil {
		log.Println("Error when inserting data to NFTStellar DB " + err.Error())
	}
	if err2 != nil {
		log.Println("Error when inserting data to MarketPlaceNFT DB " + err.Error())
	}
	return err, err2
}
