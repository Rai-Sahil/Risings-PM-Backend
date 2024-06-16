package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func StudentRoutes(r *mux.Router) {
	// GET
	r.HandleFunc("/getStudents", handlers.GetStudentHandler).Methods("GET")
	r.HandleFunc("/getStudent/{id}", handlers.GetStudentByIDHandler).Methods("GET")

	//POST
	r.HandleFunc("/addStudent", handlers.AddStudentHandler).Methods("POST")
}
