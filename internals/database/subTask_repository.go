package database

import "pm_backend/internals/models"

func InsertSubTask(subTasks models.SubTask) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Create(&subTasks).Error; err != nil {
		return err
	}

	return nil
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
