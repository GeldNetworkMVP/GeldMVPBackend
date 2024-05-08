package dataTemplates

import (
	"context"
	"log"

	"github.com/GeldNetworkMVP/GeldMVPBackend/dao"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

type Connection struct {
	*dao.Connection
}

func (cd *Connection) SaveWorkflow(obj model.PlotChainData) error {
	session, err := cd.Connect()
	if err != nil {
		log.Println("Error when connecting to DB " + err.Error())
	}
	defer session.EndSession(context.TODO())
	c := session.Client().Database(dao.DbName).Collection("PlotChainWorkflows")
	_, errx := c.InsertOne(context.TODO(), obj)

	if errx != nil {
		log.Println("Error when inserting data to PlotChainWorkflows " + err.Error())
	}
	return errx
}
