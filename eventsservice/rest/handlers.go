package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/geeks/cloud-native/lib/persistance"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

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
func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {

	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to find all available events %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying encode events to JSON %s}", err)
	}

}
func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {

	event := persistance.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occured while decoding event data %s}", err)
		return
	}
	id, err := eh.dbhandler.AddEvent(event)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occured while persisting event %d %s}", id, err)
		return
	}
}
