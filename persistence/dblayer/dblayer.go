package dblayer

import "cloud-native/persistence"

type DBTYPE string

const (
	MONGODB DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)


func NewPersistenceLayer(options DBTYPE,connection string) (persistence.DatabaseHandler,error) {
	switch options {
	case MONGODB:
		//return mongo
	}
	return nil, nil
}
