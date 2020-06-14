package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cloud-native/lib/configuration"

	"github.com/cloud-native/lib/persistance/dblayer"

	"github.com/cloud-native/eventsservice/rest"
)

func main() {

	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	//RESTful API start
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
