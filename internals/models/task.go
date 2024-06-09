package models

import "time"

type Task struct {
	ID         int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title      string    `gorm:"size:255;not null" json:"title"`
	Desc       string    `gorm:"size:255;not null" json:"desc"`
	DueDate    time.Time `gorm:"type:date;not null" json:"due_date"`
	AssigneeID int64     `gorm:"not null" json:"assignee_id"`
	Assignee   User      `gorm:"foreignKey:AssigneeID;references:ID"`
	Status     string    `gorm:"size:255;not null" json:"status"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	GoalID     *int64    `gorm:"default:null" json:"goal_id"`
	Goal       Goal      `gorm:"foreignKey:GoalID;references:ID" json:"goal"`
	Watchers   []int64   `gorm:"type:json" json:"watchers"`
	Comments   []int64   `gorm:"type:json" json:"comments"`
}
