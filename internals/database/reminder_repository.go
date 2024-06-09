package database

import (
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
	today := time.Now().Format(`2006-01-02`)

	if err := db.Where("status = ? AND DATE(due_date) = ?", "Pending", today).
		Find(&reminders).Error; err != nil {
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

func GetPendingReminderDueAfterToday() ([]models.Reminder, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var reminders []models.Reminder
	today := time.Now().Format(`2006-01-02`)

	if err := db.Where("status = ? AND due_date > ?", "Pending", today).
		Find(&reminders).Error; err != nil {
		return nil, err
	}
	return reminders, nil
}
