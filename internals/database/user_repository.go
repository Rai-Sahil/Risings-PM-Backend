package database

import (
	"pm_backend/internals/models"
)

func InsertUser(user models.User) (int64, error) {
	db, err := Connect()
	if err != nil {
		return 0, err
	}

	var existingUser models.User
	err = db.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		return existingUser.ID, nil
	}

	result := db.Create(&user)
	return user.ID, result.Error
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

func GetUserDetails(userId []int64) ([]models.User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := db.Where("id IN (?)", userId).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
