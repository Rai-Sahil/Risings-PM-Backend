package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func SubTasksRoutes(r *mux.Router) {
	r.HandleFunc("/addSubTask", handlers.InsertSubTaskHandler).Methods("POST")
	r.HandleFunc("/getSubTasks/{taskId}", handlers.GetSubTasksHandler).Methods("GET")
	r.HandleFunc("/getSubTask/{subTaskId}", handlers.GetSubTaskByIdHandler).Methods("GET")

	r.HandleFunc("/updateSubTask", handlers.UpdateSubTaskHandler).Methods("POST")
}
