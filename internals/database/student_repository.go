package database

import (
	"pm_backend/internals/models"
)

func InsertStudent(student models.Student) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Create(&student).Error; err != nil {
		return err
	}
	return nil
}

func GetAllStudents() ([]models.Student, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var students []models.Student
	if err := db.Preload("Assignee").Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
