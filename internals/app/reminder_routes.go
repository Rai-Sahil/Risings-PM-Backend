package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func ReminderRoutes(r *mux.Router) {
	r.HandleFunc("/addReminder", handlers.AddReminderHandler).Methods("POST")

	r.HandleFunc("/completeReminder/{id}", handlers.CompleteReminderHandler).Methods("GET")

	r.HandleFunc("/getPendingReminderDueToday/{assignee_id}", handlers.GetPendingReminderDueTodayByAssigneeIDHandler).
		Methods("GET")

	r.HandleFunc("/getPendingReminderDueToday", handlers.GetPendingReminderDueTodayHandler).Methods("GET")

	r.HandleFunc("/getCompletedReminder", handlers.GetCompletedReminderHandler).Methods("GET")

	r.HandleFunc("/getPendingReminderDueAfterToday", handlers.GetPendingReminderDueAfterTodayHandler).
		Methods("GET")

}
