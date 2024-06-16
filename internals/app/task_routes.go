package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func TaskRoutes(r *mux.Router) {
	//POST
	r.HandleFunc("/addTask", handlers.AddTaskHandler).Methods("POST")
	r.HandleFunc("/addComment", handlers.AddCommentHandler).Methods("POST")

	// GET
	r.HandleFunc("/getTaskById/{task_id}", handlers.GetTaskByIdHandler).Methods("GET")
	r.HandleFunc("/getPendingAdminTasks", handlers.GetPendingAdminTasksHandler).Methods("GET")
	r.HandleFunc("/getTasksByGoalId/{goal_id}", handlers.GetTasksByGoalIdHandler).Methods("GET")
	r.HandleFunc("/getTasksLengthByGoalId/{goal_id}", handlers.GetTasksLengthByGoalIdHandler).Methods("GET")
	r.HandleFunc("/getCommentsByTaskId/{task_id}", handlers.GetCommentsByTaskIdHandler).Methods("GET")
}
