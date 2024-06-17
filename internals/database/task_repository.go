package database

import (
	"pm_backend/internals/models"
	"time"
)

func InsertTask(task models.Task) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func InsertComment(comment models.Comment) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func GetCommentsByTaskID(taskID int64) ([]models.Comment, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var comments []models.Comment
	if err := db.
		Where("task_id = ?", taskID).
		Preload("User").
		Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func GetTaskByID(taskID int64) (models.Task, error) {
	db, err := Connect()
	if err != nil {
		return models.Task{}, err
	}
	var task models.Task
	if err := db.Where("id = ?", taskID).Preload("Assignee").First(&task).Error; err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func GetPendingAdminTasks() ([]models.Task, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := db.Where("status = ? AND goal_id IS NULL", "Pending").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTasksLengthByGoalId(goalId int64) (int64, error) {
	db, err := Connect()
	if err != nil {
		return -1, err
	}

	var count int64
	if err := db.Model(models.Task{}).Where("goal_id = ?", goalId).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func GetTasksByGoalId(goalId int64) ([]models.Task, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := db.Where("goal_id = ?", goalId).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTasksDueThisWeek() ([]models.Task, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	startOfWeek := now.AddDate(0, 0, -int(now.Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	var tasks []models.Task
	if err := db.
		Where("due_date >= ? AND due_date <= ?",
			startOfWeek.Format("2006-01-02"),
			endOfWeek.Format("2006-01-02")).
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
