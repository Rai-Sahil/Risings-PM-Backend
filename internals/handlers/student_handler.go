package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pm_backend/internals/database"
	"pm_backend/internals/models"
)

func AddStudentHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}
	defer r.Body.Close()

	var student models.Student
	err = json.Unmarshal(body, &student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	err = database.InsertStudent(student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Student was added successfully"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	students, err := database.GetAllStudents()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
