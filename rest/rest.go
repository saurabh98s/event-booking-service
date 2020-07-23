package rest

import (
	"cloud-native/rest/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func ServeAPI() error {
	r := mux.NewRouter()
	eventsRouter:=r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods(http.MethodGet).Path("/{SearchCriteria}/{search}").HandlerFunc()
	eventsRouter.Methods(http.MethodGet).Path("").HandlerFunc(handlers.AllEventHandler)
	eventsRouter.Methods(http.MethodPost).Path("").HandlerFunc(handlers.NewEventHandler)
	return http.ListenAndServe(":8181",r)
}