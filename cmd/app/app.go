package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"pm_backend/internals/app"
)

func main() {
	r := mux.NewRouter()

	app.UserRoutes(r)
	app.StudentRoutes(r)
	app.GoalRoutes(r)
	app.ReminderRoutes(r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
