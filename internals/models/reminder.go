package models

import "time"

type Reminder struct {
	ID         int64     `gorm:"primary_key;auto-increment" json:"id"`
	Title      string    `gorm:"size:255;not null" json:"title"`
	Desc       string    `gorm:"size:255;not null" json:"desc"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	DueDate    time.Time `gorm:"autoCreateTime" json:"due_date"`
	Priority   string    `gorm:"size:255;default:Low" json:"priority"`
	Status     string    `gorm:"size:255;default:Pending" json:"status"`
	AssigneeID int64     `gorm:"not null" json:"assignee_id"`
	Assignee   User      `gorm:"foreignKey:AssigneeID;references:ID" json:"assignee"`
}
