package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"pm_backend/internals/app"
)

func main() {
	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	app.UserRoutes(router)
	app.StudentRoutes(router)
	app.GoalRoutes(router)
	app.ReminderRoutes(router)
	app.TaskRoutes(router)
	app.AuthRoutes(router)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
