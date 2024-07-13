package database

import (
	"pm_backend/internals/models"
)

func InsertGoal(goal models.Goal) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Create(&goal).Error; err != nil {
		return err
	}
	return nil
}

func GetPendingGoalsCountByUserId(userId int64) (int64, error) {
	db, err := Connect()
	if err != nil {
		return -1, err
	}
	var count int64
	if err := db.
		Where("assignee_id = ? AND status != ?", userId, "Completed").
		Model(&models.Goal{}).
		Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func GetInCompletedGoalsCountByStudentId(studentId int64) (int64, error) {
	db, err := Connect()
	if err != nil {
		return -1, err
	}
	var count int64
	if err := db.
		Where("student_id = ? AND status = ?", studentId, "Pending").
		Model(&models.Goal{}).
		Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func GetAllGoals() ([]models.Goal, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var goals []models.Goal
	if err := db.Preload("Assignee").Preload("Student").Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}

func GetGoalsByStudentID(studentID int64) ([]models.Goal, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var goals []models.Goal
	if err = db.Where("student_id = ?", studentID).Preload("Assignee").Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}

func GetGoalByGoalID(goalID int64) ([]models.Goal, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var goals []models.Goal
	if err = db.Where("id = ?", goalID).
		Preload("Student").
		Preload("Assignee").
		Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}

func GetPendingGoalsByStudentID(studentID int64) ([]models.Goal, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var goals []models.Goal
	if err := db.Where("student_id = ?", studentID).
		Where("status = ?", "Pending").
		Preload("Assignee").
		Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}

func GetPendingGoalsByUserId(userIds []int64) ([]models.Goal, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var goals []models.Goal
	if err := db.
		Where("assignee_id IN (?) and status != ?", userIds, "Completed").
		Preload("Assignee").
		Find(&goals).Error; err != nil {
		return nil, err
	}

	return goals, nil
}

func GetInCompleteGoals() ([]models.Goal, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	var goals []models.Goal
	if err := db.Where("status != ?", "Complete").Find(&goals).Error; err != nil {
		return nil, err
	}

	return goals, nil
}

func UpdateGoal(goal models.Goal) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.Save(&goal).Error; err != nil {
		return err
	}

	return nil
}
