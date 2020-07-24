package main

import (
	"cloud-native/configuration"
	"cloud-native/logger"
	"cloud-native/persistence/dblayer"
	"cloud-native/rest"
	"flag"
	"log"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	logger.Log.Info("[DEBUG] Connecting to database")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	logger.Log.Info("[DEBUG] Connected to database")
	//REST API start
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
