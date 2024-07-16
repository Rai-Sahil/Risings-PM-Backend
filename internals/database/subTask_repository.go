package database

import (
	"pm_backend/internals/models"
)

func InsertSubTask(subTasks models.SubTask) (models.SubTask, error) {
	db, err := Connect()
	if err != nil {
		return subTasks, err
	}

	if err := db.Create(&subTasks).Error; err != nil {
		return subTasks, err
	}

	if err := db.Preload("Assignee").First(&subTasks, subTasks.ID).Error; err != nil {
		return subTasks, err
	}

	return subTasks, nil
}

func GetSubTasks(taskId int64) ([]models.SubTask, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var subTasks []models.SubTask
	if err := db.Where("task_id = ?", taskId).Preload("Assignee").Find(&subTasks).Error; err != nil {
		return nil, err
	}

	return subTasks, nil
}

func GetSubTaskById(subTaskId int64) (models.SubTask, error) {
	db, err := Connect()
	if err != nil {
		return models.SubTask{}, err
	}

	var subTask models.SubTask
	if err := db.Where("id = ?", subTaskId).Preload("Assignee").Preload("Task").First(&subTask).Error; err != nil {
		return models.SubTask{}, err
	}

	return subTask, nil
}

func UpdateSubTask(subTask models.SubTask) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Save(&subTask).Error; err != nil {
		return err
	}

	return nil
}

func DeleteSubTask(subTaskID int64) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	tx := db.Begin()

	if err := tx.Where("id = ?", subTaskID).Delete(&models.SubTask{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func GetCommentsBySubTaskID(subTaskID int64) ([]models.Comment, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var comments []models.Comment
	if err := db.
		Where("sub_task_id = ?", subTaskID).
		Preload("User").
		Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}
