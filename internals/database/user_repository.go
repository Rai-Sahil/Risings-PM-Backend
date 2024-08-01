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

func SetUserStatus(userId int64, status string) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	var user models.User
	err = db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return err
	}

	user.Status = status

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
