package main

import (
	"cloud-native/rest"
	"log"
)

func main() {
	log.Fatal(rest.ServeAPI())
}
