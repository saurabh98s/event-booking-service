package rest

import (
	"cloud-native/persistence"
	"cloud-native/rest/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// ServeAPI handles the endpoints and db type to serve
func ServeAPI(endpoint string,tlsendpoint string,databasehandler persistence.DatabaseHandler)(chan error,chan error) {
	handler := handlers.NewEventHandler(databasehandler)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.AddEventHandler)
	httpErrChan := make(chan error)
	httptlsErrChan := make(chan error)

	go func() {
		httptlsErrChan <- http.ListenAndServeTLS(tlsendpoint, "cert.pem", "key.pem", r)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(endpoint, r)
	}()

	return httpErrChan, httptlsErrChan
}
