package dblayer

import (
	"cloud-native/persistence"
	"cloud-native/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)


func NewPersistenceLayer(options DBTYPE,connection string) (persistence.DatabaseHandler,error) {
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}