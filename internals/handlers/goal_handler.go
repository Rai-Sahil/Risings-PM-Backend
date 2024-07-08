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

func AddGoalHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var goal models.Goal
	err = json.Unmarshal(body, &goal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.InsertGoal(goal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Goal was added successfully"))
}

func GetPendingGoalsCountByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userIdStr := mux.Vars(r)["user_id"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	count, err := database.GetPendingGoalsCountByUserId(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if err := json.NewEncoder(w).Encode(count); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func GetGoalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	goal, err := database.GetAllGoals()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(goal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetPendingGoalsByStudentIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	studentIDStr := mux.Vars(r)["student_id"]
	studentID, err := strconv.ParseInt(studentIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid student ID"))
	}

	goals, err := database.GetPendingGoalsByStudentID(studentID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if err := json.NewEncoder(w).Encode(goals); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func GetGoalByStudentIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	studentIDStr := mux.Vars(r)["student_id"]
	studentID, err := strconv.ParseInt(studentIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid student ID"))
	}

	goals, err := database.GetGoalsByStudentID(studentID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Database error while getting goals"))
	}
	if err := json.NewEncoder(w).Encode(goals); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func GetInCompletedGoalsCountByStudentIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	studentIDStr := mux.Vars(r)["student_id"]
	studentID, err := strconv.ParseInt(studentIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid student ID"))
	}

	count, err := database.GetInCompletedGoalsCountByStudentId(studentID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	if err := json.NewEncoder(w).Encode(count); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func GetGoalByGoalIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	goalIDStr := mux.Vars(r)["goal_id"]
	goalID, err := strconv.ParseInt(goalIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid goal ID"))
	}

	goals, err := database.GetGoalByGoalID(goalID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Database error while getting goals"))
	}
	if err := json.NewEncoder(w).Encode(goals); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Database error while getting goals"))
	}
	w.WriteHeader(http.StatusOK)
}

func GetPendingGoalsByUserIdHandlers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		UserIds []int64 `json:"userIds"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	goals, err := database.GetPendingGoalsByUserId(request.UserIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(goals); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateGoalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var goal models.Goal

	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := database.UpdateGoal(goal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(goal); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}

func GetInCompleteGoalsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	goals, err := database.GetInCompleteGoals()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if err := json.NewEncoder(w).Encode(goals); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
