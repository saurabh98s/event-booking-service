package handlers

import (
	"log"
	"net/http"
)

type eventServiceHandler struct {

}

func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter,r *http.Request) {
		log.Println("FInd event Handler")
}

func (eh eventServiceHandler) AllEventHandler(w http.ResponseWriter,r *http.Request) {
	log.Println("All event Handler")
}
func (eh eventServiceHandler) NewEventHandler(w http.ResponseWriter,r *http.Request) {
	log.Println("New event Handler")
}
