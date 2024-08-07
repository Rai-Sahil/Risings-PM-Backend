package app

import (
	"github.com/gorilla/mux"
	"pm_backend/internals/handlers"
)

func ReminderRoutes(r *mux.Router) {
	r.HandleFunc("/addReminder", handlers.AddReminderHandler).Methods("POST")

	r.HandleFunc("/getReminder/{reminderId}", handlers.GetReminderHandler).Methods("GET")
	r.HandleFunc("/deleteReminder/{reminderId}", handlers.DeleteReminderHandler).Methods("DELETE")

	r.HandleFunc("/updateReminder", handlers.UpdateReminderHandler).Methods("POST")

	r.HandleFunc("/completeReminder/{id}", handlers.CompleteReminderHandler).Methods("GET")

	r.HandleFunc("/getPendingReminderDueToday/{assignee_id}", handlers.GetPendingReminderDueTodayByAssigneeIDHandler).
		Methods("GET")

	r.HandleFunc("/getPendingReminderDueToday", handlers.GetPendingReminderDueTodayHandler).Methods("GET")

	r.HandleFunc("/getCompletedReminder", handlers.GetCompletedReminderHandler).Methods("GET")

	r.HandleFunc("/getPendingReminderDueAfterToday", handlers.GetPendingReminderDueAfterTodayHandler).
		Methods("GET")

	r.HandleFunc("/getPendingReminderCountDueToday/{assignee_id}", handlers.GetPendingReminderDueTodayCountByUserIdHandler).
		Methods("GET")

	r.HandleFunc("/getPendingReminderDueTodayMultipleUsers", handlers.GetPendingReminderDueTodayByMultipleUsersHandler).Methods("POST")
}
