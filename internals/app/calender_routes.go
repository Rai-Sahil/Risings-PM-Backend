package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func CalenderRoutes(r *mux.Router) {
	r.HandleFunc("/addEvent", handlers.InsertEventHandler).Methods("POST")
	r.HandleFunc("/getEvents", handlers.GetEventsHandler).Methods("GET")
}
