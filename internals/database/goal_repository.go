package database

import "pm_backend/internals/models"

func InsertGoal(goal models.Goal) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	result := db.Create(&goal)
	return result.Error
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
	if err = db.Where("goal_id = ?", goalID).
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
	if err = db.Where("student_id = ?", studentID).
		Not("status = ?", "Pending").
		Preload("Assignee").
		Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}
