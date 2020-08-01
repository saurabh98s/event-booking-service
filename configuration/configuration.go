package configuration

import (
	"cloud-native/persistence/dblayer"
	"encoding/json"
	"fmt"
	"os"
)

var (
	//DBTypeDefault holds the type of database to be used
	DBTypeDefault           = dblayer.DBTYPE("mongodb")
	// DBConnectionDefault holds the default connection string
	DBConnectionDefault     = "mongodb://127.0.0.1:27017"
	// RestfulEndPointPDefault holds the http endpoint
	RestfulEndPointPDefault = "localhost:8181"
	// RestfulTLSEndPointPDefault holds the https endpoint
	RestfulTLSEndPointPDefault = "localhost:9191"
)

type ServiceConfig struct {
	Databasetype    dblayer.DBTYPE `json:"databasetype"`
	DBConnection    string         `json:"dbconnection"`
	RestfulEndpoint string         `json:"restfulapi_endpoint"`
	RestfulTLSEndpoint string       `json:"restfulapi_tlsendpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEndPointPDefault,
		RestfulTLSEndPointPDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
