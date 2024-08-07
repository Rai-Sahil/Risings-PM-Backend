package database

import (
	"gorm.io/gorm"
	"pm_backend/internals/models"
	"time"
)

func InsertTask(task models.Task) (models.Task, error) {
	db, err := Connect()
	if err != nil {
		return task, err
	}

	if err := db.Create(&task).Error; err != nil {
		return task, err
	}

	if err := db.Preload("Assignee").First(&task, task.ID).Error; err != nil {
		return task, err
	}

	return task, nil
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

func InsertComment(comment models.Comment) (models.Comment, error) {
	db, err := Connect()
	if err != nil {
		return models.Comment{}, err
	}

	if err := db.Create(&comment).Error; err != nil {
		return models.Comment{}, err
	}

	if err := db.Preload("Task").Preload("SubTask").Preload("User").First(&comment, comment.ID).Error; err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func UpdateComment(id int64, content string) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Model(&models.Comment{}).Where("id = ?", id).Update("content", content).Error; err != nil {
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

func GetAllAdminTasks() ([]models.Task, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := db.Where("goal_id IS NULL").Preload("Assignee").Find(&tasks).Error; err != nil {
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
	sevenDaysLater := now.AddDate(0, 0, 7)
	var tasks []models.Task

	if err := db.
		Where("due_date BETWEEN ? AND ?", now, sevenDaysLater).
		Preload("Assignee").
		Preload("Goal").Preload("Goal.Student").
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

func GetPendingTasksDueTodayByUserId(userId int64) ([]models.Task, error) {
	db, err := Connect()
	if err != nil {
		return []models.Task{}, err
	}

	var tasks []models.Task

	err = db.Model(models.Task{}).
		Where("assignee_id = ? AND status != ?", userId, "Completed").
		Preload("Assignee").
		Preload("Goal").Preload("Goal.Student").
		Find(&tasks).Error

	if err != nil {
		return []models.Task{}, err
	}

	return tasks, nil
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

func GetUsersUpdateForToday() ([]map[string]interface{}, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	today := time.Now().Format("2006-01-02")

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	var userTaskCounts []map[string]interface{}

	for _, user := range users {
		var dueCount int64
		if err := db.Model(&models.Task{}).
			Where("due_date = ? AND assignee_id = ?", today, user.ID).
			Count(&dueCount).Error; err != nil {
			return nil, err
		}

		var completedCount int64
		if err := db.Model(&models.Task{}).
			Where("due_date = ? AND assignee_id = ? AND status = ?", today, user.ID, "Completed").
			Count(&completedCount).Error; err != nil {
			return nil, err
		}

		userTaskCounts = append(userTaskCounts, map[string]interface{}{
			"name":      user.Name,
			"due":       int(dueCount),
			"completed": int(completedCount),
		})
	}

	return userTaskCounts, nil
}
