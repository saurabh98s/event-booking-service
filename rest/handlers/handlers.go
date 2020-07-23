package handlers

import (
	"cloud-native/persistence"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

func NewEventHandler(databasehandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler: databasehandler,
	}
}

func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter,r *http.Request) {
		vars:=mux.Vars(r)
		criteria,ok:=vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w,`{"error":"No search criteria found, you can either 
								search by id via /id/4 
								to search by name via /name/coldplayconcert"`)
		return
	}
	searchKey,ok:=vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w,`{"error":"No search criteria found, you can either 
								search by id via /id/4 
								to search by name via /name/coldplayconcert"`)
		return

	}
}

func (eh eventServiceHandler) AllEventHandler(w http.ResponseWriter,r *http.Request) {
	log.Println("All event Handler")
}
func (eh eventServiceHandler) NewEventHandler(w http.ResponseWriter,r *http.Request) {
	log.Println("New event Handler")
}
