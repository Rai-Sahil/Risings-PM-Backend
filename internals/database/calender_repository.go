package database

import "pm_backend/internals/models"

func InsertEvent(calender models.Calendar) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Create(&calender).Error; err != nil {
		return err
	}
	return nil
}

func GetEvents() ([]models.Calendar, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var calendars []models.Calendar
	if err := db.Preload("Assignee").Find(&calendars).Error; err != nil {
		return nil, err
	}

	return calendars, nil
}
