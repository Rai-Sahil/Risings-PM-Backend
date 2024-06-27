package database

import (
	"errors"
	"pm_backend/internals/models"
	"time"
)

func InsertReminder(reminder models.Reminder) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Create(&reminder).Error; err != nil {
		return err
	}
	return nil
}

func GetPendingReminderDueTodayByAssigneeID(assigneeID int64) ([]models.Reminder, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var reminders []models.Reminder
	today := time.Now().Format(`2006-01-02`)

	if err := db.Where("assignee_id = ? AND status = ? AND DATE(due_date) = ?", assigneeID, "Pending", today).
		Find(&reminders).Error; err != nil {
		return nil, err
	}
	return reminders, nil
}

func GetPendingReminderDueToday() ([]models.Reminder, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var reminders []models.Reminder
	today := time.Now().Format("2006-01-02")

	if err := db.Where("status = ? AND DATE(due_date) = ?", "Pending", today).Find(&reminders).Error; err != nil {
		return nil, err
	}
	return reminders, nil
}

func GetCompletedReminder() ([]models.Reminder, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var reminders []models.Reminder
	if err := db.Where("status = ?", "Completed").Find(&reminders).Error; err != nil {
		return nil, err
	}
	return reminders, nil
}

func CompleteReminder(id int64) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	// Find the reminder by ID
	var reminder models.Reminder
	if err := db.First(&reminder, id).Error; err != nil {
		return err
	}

	// Check if the reminder is already complete
	if reminder.Status == "Complete" {
		return errors.New("reminder is already marked as complete")
	}

	// Update the status to complete
	reminder.Status = "Complete"
	if err := db.Save(&reminder).Error; err != nil {
		return err
	}

	return nil
}

func GetPendingReminderDueAfterToday() ([]models.Reminder, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var reminders []models.Reminder
	today := time.Now().Format("2006-01-02")

	if err := db.Where("status = ? AND DATE(due_date) > ?", "Pending", today).Find(&reminders).Error; err != nil {
		return nil, err
	}
	return reminders, nil
}

func GetPendingReminderDueTodayCountByUserId(userID int64) (int64, error) {
	db, err := Connect()
	if err != nil {
		return 0, err
	}

	var count int64
	today := time.Now().Format("2006-01-02")

	err = db.Model(&models.Reminder{}).
		Where("assignee_id = ? AND status = ? AND DATE(due_date) = ?", userID, "Pending", today).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
