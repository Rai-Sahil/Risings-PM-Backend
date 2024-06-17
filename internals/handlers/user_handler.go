package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pm_backend/internals/database"
	"pm_backend/internals/models"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var newUser models.User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.InsertUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintf(w, "User added successfully")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := database.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
