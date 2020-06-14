package dblayer

import (
	"github.com/geeks/cloud-native/lib/persistance"
	persistance2 "github.com/geeks/cloud-native/lib/persistance/mongolayer"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistance.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return persistance2.NewMongoDBLayer(connection)
	}
	return nil, nil
}
