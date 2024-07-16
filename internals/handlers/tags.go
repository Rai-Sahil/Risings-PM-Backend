package handlers

import (
	"encoding/json"
	"net/http"
	"pm_backend/internals/database"
)

func GetTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tags := database.NewGoalTagManager.GetTags()
	json.NewEncoder(w).Encode(tags)
}

func AddGoalTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tag database.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.NewGoalTagManager.AddTag(tag.Name, tag.Color); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tag)
}

func DeleteGoalTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing tag name", http.StatusBadRequest)
		return
	}

	if err := database.NewGoalTagManager.DeleteTag(name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ChangeGoalTagName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req struct {
		OldName string `json:"old_name"`
		NewName string `json:"new_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.NewGoalTagManager.ChangeTagName(req.OldName, req.NewName); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
