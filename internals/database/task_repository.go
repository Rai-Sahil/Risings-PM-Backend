package database

import (
	"gorm.io/gorm"
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

func UpdateTask(task models.Task) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Save(&task).Error; err != nil {
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
	if err := db.Where("id = ?", taskID).Preload("Goal").Preload("Assignee").First(&task).Error; err != nil {
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
	if err := db.Where("status = ? AND goal_id IS NULL", "Pending").Preload("Assignee").Find(&tasks).Error; err != nil {
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
	if err := db.Where("goal_id = ?", goalId).Preload("Assignee").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTasksByUserIdAndGoalId(userIds []int64, goalId int64) ([]models.Task, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := db.Where("goal_id = ? AND assignee_id IN (?)", goalId, userIds).
		Preload("Assignee").
		Order(gorm.Expr("CASE WHEN status = 'Overdue' THEN 1 WHEN status = 'In Progress' THEN 2 WHEN status = 'Pending' THEN 3 ELSE 4 END")).
		Find(&tasks).Error; err != nil {
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
		Preload("Assignee").
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTotalTasksByUserIdThisWeekCount(userId int64) (int64, error) {
	db, err := Connect()
	if err != nil {
		return -1, err
	}

	now := time.Now()

	offset := (int(time.Sunday) - int(now.Weekday()) - 7) % 7
	startOfWeek := now.AddDate(0, 0, offset)
	startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, startOfWeek.Location())

	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	var count int64
	if err := db.Model(models.Task{}).
		Where("assignee_id = ? AND due_date BETWEEN ? AND ?", userId, startOfWeek, endOfWeek).
		Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func GetTasksCompleteByUserIdThisWeekCount(userId int64) (int64, error) {
	db, err := Connect()
	if err != nil {
		return -1, err
	}

	// Get the current time
	now := time.Now()

	// Calculate the start of the week (assuming week starts on Sunday)
	offset := (int(time.Sunday) - int(now.Weekday()) - 7) % 7
	startOfWeek := now.AddDate(0, 0, offset)
	startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, startOfWeek.Location())

	// Calculate the end of the week
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	var count int64
	if err := db.Model(models.Task{}).
		Where("assignee_id = ? AND status = ? AND created_at BETWEEN ? AND ?", userId, "Completed", startOfWeek, endOfWeek).
		Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func GetPendingTasksCountDueTodayByUserId(userId int64) (int64, error) {
	db, err := Connect()
	if err != nil {
		return -1, err
	}

	now := time.Now().Format("2006-01-02")
	var count int64

	err = db.Model(models.Task{}).
		Where("assignee_id = ? AND due_date = ? AND status = ?", userId, now, "Pending").
		Count(&count).Error

	if err != nil {
		return -1, err
	}

	return count, nil
}

func GetCompletedTasksCountDueTodayByUserId(userId int64) (int64, error) {
	db, err := Connect()
	if err != nil {
		return -1, err
	}

	now := time.Now().Format("2006-01-02")
	var count int64

	err = db.Model(models.Task{}).
		Where("assignee_id = ? AND due_date = ? AND status = ?", userId, now, "Completed").
		Count(&count).Error

	if err != nil {
		return -1, err
	}

	return count, nil
}

func GetTasksDueThisWeekByUserIds(userIds []int64) ([]models.Task, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	sevenDaysLater := now.AddDate(0, 0, 7)
	var tasks []models.Task

	if err := db.Where("assignee_id IN (?) AND due_date BETWEEN ? AND ?", userIds, now, sevenDaysLater).
		Preload("Assignee").
		Order("due_date ASC").
		Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTasksCountGroupByStatusByGoalId(goalId int64) (map[string]int64, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	type StatusCount struct {
		Status string `json:"status"`
		Count  int64  `json:"count"`
	}

	var statusCounts []StatusCount
	if err := db.
		Model(&models.Task{}).
		Select("status, COUNT(*) as count").
		Where("goal_id = ?", goalId).
		Group("status").
		Scan(&statusCounts).Error; err != nil {
		return nil, err
	}

	countMap := make(map[string]int64)
	for _, count := range statusCounts {
		countMap[count.Status] = count.Count
	}

	return countMap, nil
}

func GetOverDueIncompleteTasksCountByGoalId(goalId int64) (int64, error) {
	db, err := Connect()
	if err != nil {
		return -1, err
	}

	var count int64
	if err := db.
		Model(&models.Task{}).
		Where("status != ? AND goal_id = ? AND due_date < ?", "Completed", goalId, time.Now()).
		Count(&count).Error; err != nil {
		return -1, err
	}

	return count, nil
}

func DeleteTask(taskID int64) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	tx := db.Begin()

	if err := tx.Where("task_id = ?", taskID).Delete(&models.SubTask{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", taskID).Delete(&models.Task{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
