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

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var task models.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = database.InsertTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Task added"))
}

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	defer r.Body.Close()

	var comment models.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	err = database.InsertComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
}

func GetCommentsByTaskIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskIdStr := mux.Vars(r)["task_id"]
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	comments, err := database.GetCommentsByTaskID(taskId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func GetPendingAdminTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := database.GetPendingAdminTasks()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskIdStr := mux.Vars(r)["task_id"]
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	task, err := database.GetTaskByID(taskId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func GetTasksLengthByGoalIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	goalIdStr := mux.Vars(r)["goal_id"]
	goalId, err := strconv.ParseInt(goalIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	count, err := database.GetTasksLengthByGoalId(goalId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(count); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func GetTasksByGoalIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	goalIdStr := mux.Vars(r)["goal_id"]
	goalId, err := strconv.ParseInt(goalIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	tasks, err := database.GetTasksByGoalId(goalId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
