package rest

import (
	"cloud-native/persistence"
	"cloud-native/rest/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func ServeAPI(endpoint string,databasehandler persistence.DatabaseHandler) error {
	handler:=NewEventHandler(databasehandler)
	r := mux.NewRouter()
	eventsRouter:=r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods(http.MethodGet).Path("/{SearchCriteria}/{search}").HandlerFunc(handlers.)
	eventsRouter.Methods(http.MethodGet).Path("").HandlerFunc(handlers.AllEventHandler)
	eventsRouter.Methods(http.MethodPost).Path("").HandlerFunc(handlers.NewEventHandler)
	return http.ListenAndServe(":8181",r)
}