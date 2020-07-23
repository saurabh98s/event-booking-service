package main

import(
	"flag"
	"fmt"
	"github.com/cloud-native/eventsservice/rest"
	"github.com/cloud-native/lib/configuration"
	"github.com/cloud-native/lib/persistance/dblayer"
	"log"
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
