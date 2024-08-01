package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func UserRoutes(r *mux.Router) {

	r.HandleFunc("/getUsers", handlers.GetUserHandler).Methods("GET")

	r.HandleFunc("/addUser", handlers.AddUserHandler).Methods("POST")
	r.HandleFunc("/setUserStatus/{user_id}/{status}", handlers.SetUserStatusHandler).Methods("GET")
}
