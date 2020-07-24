package handlers

import (
	"cloud-native/logger"
	"cloud-native/persistence"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

func NewEventHandler(databaseHandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler: databaseHandler,
	}
}

func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error":"No search criteria found, you can either 
								search by id via /id/4 
								to search by name via /name/coldplayconcert"`)
		logger.Log.Error("No search criteria found, you can either search by id via /id/4to search by name via /name/coldplayconcert")
		return
	}
	searchKey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error":"No search criteria found, you can either 
								search by id via /id/4 
								to search by name via /name/coldplayconcert"`)
		logger.Log.Error("No search criteria found, you can either search by id via /id/4to search by name via /name/coldplayconcert")
		return

	}

	var event persistence.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchKey)
	case "id":
		id, err := hex.DecodeString(searchKey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)

}

func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil || events == nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying to find all available events %s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying encode events to JSON %s"}`, err)
	}
}
func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	logger.Log.WithTime(time.Now().Local()).Info("[DEBUG] Adding New Event to DB" ,event.Name)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "error occured while decoding event data %s"}`, err)
		return
	}
	result, err := eh.dbhandler.AddEvent(event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "error occured while persisting event %s:  %s"}`, result.ID.Hex(), err)
		return
	}
	fmt.Fprintf(w, `{"id":%d}`, result.ID.Hex())
	logger.Log.Info("[DEBUG] Added event: ",result.Name)
}
