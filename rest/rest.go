package rest

import (
	"cloud-native/persistence"
	"cloud-native/rest/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// ServeAPI handles the endpoints and db type to serve
func ServeAPI(endpoint string, databasehandler persistence.DatabaseHandler) error {
	handler := handlers.NewEventHandler(databasehandler)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(endpoint, r)
}
