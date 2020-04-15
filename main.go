package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"persistance"
	"strings"

	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string) error {
	handler := &eventServiceHandler{}
	router := mux.NewRouter()
	/* here now we will create an api interface
	which will be able to
	1. Search events
	2. Retrieve all events
	3. Create new events
	*/
	eventsRouter := router.PathPrefix("/events").Subrouter()
	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	return http.ListenAndServe(":8181", router)
}

type eventServiceHandler struct {
	dbhandler persistance.DatabaseHandler
}

func newEventHandler(databasehandler persistance.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler: databasehandler,
	}
}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search criteria found, you can either search by id via /id/4
				   to search by name via /name/coldplayconcert}`)
		return
	}
	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search keys found, you can either search by id via /id/4
						to search by name via /name/coldplayconcert}`)
		return
	}
	var event persistance.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchkey)
	case "id":
		id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}
	if err != nil {
		fmt.Fprintf(w, "{error %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}
func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {}
func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {}
