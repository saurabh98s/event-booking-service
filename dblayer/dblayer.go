package dblayer

import (
	//"github.com/geeks/cloud-native/persistance"
	"github.com/geeks/cloud-native/persistance/mongolayer"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistance.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return persistance.NewMongoDBLayer(connection)
	}
	return nil, nil
}
