package models

import "time"

type Goal struct {
	ID             int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title          string    `gorm:"size:255;not null" json:"title"`
	Desc           string    `gorm:"size:255" json:"desc"`
	Status         string    `gorm:"size:255;not null;default:Pending" json:"status"`
	StudentID      int64     `gorm:"not null" json:"student_id"`
	Student        Student   `gorm:"foreignKey:StudentID;references:ID" json:"student"`
	AssigneeID     int64     `gorm:"not null" json:"assignee_id"`
	Assignee       User      `gorm:"foreignKey:AssigneeID;references:ID" json:"assignee"`
	NumberOfTasks  int64     `gorm:"not null;default:0" json:"number_of_tasks"`
	TasksCompleted int64     `gorm:"not null;default:0" json:"tasks_complete"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	StartDate      time.Time `gorm:"default:null" json:"start_date"`
	EndDate        time.Time `gorm:"default:null" json:"end_date"`
	Priority       string    `gorm:"default:Low" json:"priority"`
	Tasks          []Task    `gorm:"foreignKey:GoalID;constraint:OnDelete:CASCADE;" json:"tasks"`
}
