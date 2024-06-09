package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func UserRoutes(r *mux.Router) {

	r.HandleFunc("/addUser", handlers.AddUserHandler).Methods("POST")
	r.HandleFunc("/getUsers", handlers.GetUserHandler).Methods("GET")
}
