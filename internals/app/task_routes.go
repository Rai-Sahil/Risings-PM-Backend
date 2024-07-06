package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func TaskRoutes(r *mux.Router) {
	//POST
	r.HandleFunc("/addTask", handlers.AddTaskHandler).Methods("POST")
	r.HandleFunc("/addComment", handlers.AddCommentHandler).Methods("POST")
	r.HandleFunc("/getTaskByGoalAndUserId", handlers.GetTasksByUserIdAndGoalIdHandler).Methods("POST")
	r.HandleFunc("/getTasksDueThisWeekByUserId", handlers.GetTasksDueThisWeekByUserIdsHandler).Methods("POST")
	r.HandleFunc("/updateTask", handlers.UpdateTaskHandler).Methods("POST")

	// GET
	r.HandleFunc("/getTasksDueThisWeek", handlers.GetTasksDueThisWeekHandler).Methods("GET")
	r.HandleFunc("/getTaskById/{task_id}", handlers.GetTaskByIdHandler).Methods("GET")
	r.HandleFunc("/getPendingAdminTasks", handlers.GetPendingAdminTasksHandler).Methods("GET")
	r.HandleFunc("/getTasksByGoalId/{goal_id}", handlers.GetTasksByGoalIdHandler).Methods("GET")
	r.HandleFunc("/getTasksLengthByGoalId/{goal_id}", handlers.GetTasksLengthByGoalIdHandler).Methods("GET")
	r.HandleFunc("/getCommentsByTaskId/{task_id}", handlers.GetCommentsByTaskIdHandler).Methods("GET")
	r.HandleFunc("/getTasksCountByUserIdThisWeek/{user_id}", handlers.GetTotalTasksByUserIdThisWeekCountHandler).
		Methods("GET")
	r.HandleFunc("/getCompletedTasksByUserId/{user_id}", handlers.GetTasksCompleteByUserIdThisWeekHandler).
		Methods("GET")

	r.HandleFunc("/getPendingTasksDueToday/{user_id}", handlers.GetPendingTasksDueTodayCountHandler).Methods("GET")
	r.HandleFunc("/getCompletedTasksDueToday/{user_id}", handlers.GetCompletedTasksDueTodayCountHandler).Methods("GET")
}
