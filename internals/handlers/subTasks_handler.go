package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"pm_backend/internals/database"
	"pm_backend/internals/models"
	"strconv"
)

func InsertSubTaskHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var subTask models.SubTask
	err = json.Unmarshal(body, &subTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	err = database.InsertSubTask(subTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
}

func GetSubTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskIdStr := mux.Vars(r)["taskId"]
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid student ID"))
	}

	subtasks, err := database.GetSubTasks(taskId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if err := json.NewEncoder(w).Encode(subtasks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
