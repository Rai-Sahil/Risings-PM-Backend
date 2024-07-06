package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pm_backend/internals/database"
	"pm_backend/internals/models"
)

func InsertEventHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var calenderEvent models.Calendar
	err = json.Unmarshal(body, &calenderEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	err = database.InsertEvent(calenderEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	calenderEvents, err := database.GetEvents()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(calenderEvents); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
