package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func GoalTagRoutes(r *mux.Router) {
	r.HandleFunc("/getGoalTags", handlers.GetTags).Methods("GET")
	r.HandleFunc("/deleteGoalTag", handlers.DeleteGoalTag).Methods("DELETE")
	r.HandleFunc("/changeGoalTagName", handlers.ChangeGoalTagName).Methods("PUT")
	r.HandleFunc("/addGoalTag", handlers.AddGoalTag).Methods("POST")
}
