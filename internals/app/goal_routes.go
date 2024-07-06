package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func GoalRoutes(r *mux.Router) {
	r.HandleFunc("/addGoal", handlers.AddGoalHandler).Methods("POST")
	r.HandleFunc("/getPendingGoalsByMultipleUserID", handlers.GetPendingGoalsByUserIdHandlers).Methods("POST")

	r.HandleFunc("/getPendingGoalByUserID/{user_id}", handlers.GetPendingGoalsCountByUserIDHandler).Methods("GET")
	r.HandleFunc("/getGoals", handlers.GetGoalHandler).Methods("GET")
	r.HandleFunc("/getGoalsByStudentID/{student_id}", handlers.GetGoalByStudentIDHandler).Methods("GET")
	r.HandleFunc("/getGoals/{goal_id}", handlers.GetGoalByGoalIDHandler).Methods("GET")
	r.HandleFunc("/getPendingGoalsByStudentID/{student_id}", handlers.GetPendingGoalsByStudentIDHandler).Methods("GET")

}
