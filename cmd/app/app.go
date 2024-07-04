package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"os"
	"pm_backend/internals/app"
)

func main() {
	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://pm-frontend-swart.vercel.app", "http://54.218.199.46", "http://54.218.199.46:80"},
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		panic(err)
	}
}
