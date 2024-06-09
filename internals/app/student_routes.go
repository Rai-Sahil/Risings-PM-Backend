package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func StudentRoutes(r *mux.Router) {
	r.HandleFunc("/addStudent", handlers.AddStudentHandler).Methods("POST")
	r.HandleFunc("/getStudents", handlers.GetStudentHandler).Methods("GET")
}
