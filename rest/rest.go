package rest

import (
	"github.com/geeks/cloud-native/persistance"
	"net/http"

	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string, dbHandler persistance.DatabaseHandler) error {
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
