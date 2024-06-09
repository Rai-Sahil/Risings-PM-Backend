package database

import "pm_backend/internals/models"

func InsertUser(user models.User) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	result := db.Create(&user)
	return result.Error
}

func GetAllUsers() ([]models.User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var users []models.User
	db.Find(&users)
	return users, nil
}
