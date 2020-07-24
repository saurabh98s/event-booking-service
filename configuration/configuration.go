package configuration

import (
	"cloud-native/persistence/dblayer"
	"encoding/json"
	"fmt"
	"os"
)

var (
	DBTypeDefault           = dblayer.DBTYPE("mongodb")
	DBConnectionDefault     = "mongodb://127.0.0.1:27017"
	RestfulEndPointPDefault = "localhost:8181"
)

type ServiceConfig struct {
	Databasetype    dblayer.DBTYPE `json:"databasetype"`
	DBConnection    string         `json:"dbconnection"`
	RestfulEndpoint string         `json:"restfulapi_endpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEndPointPDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
