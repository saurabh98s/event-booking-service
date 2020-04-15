package dblayer

import (
	"github.com/cloud-native/lib/persistance"
	// . "github.com/cloud-native/persistance"
	mongolayer "github.com/cloud-native/lib/persistance/mongolayer"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistance.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
