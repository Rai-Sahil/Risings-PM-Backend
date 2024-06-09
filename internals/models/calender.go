package models

import "time"

type Calendar struct {
	ID         int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title      string    `gorm:"size:255;not null" json:"title"`
	Desc       string    `gorm:"size:255;not null" json:"desc"`
	StartDate  time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate    time.Time `gorm:"type:date;not null" json:"end_date"`
	AssigneeID int64     `gorm:"not null" json:"assignee_id"`
	Assignee   User      `gorm:"foreignKey:AssigneeID;references:ID"`
	Status     string    `gorm:"size:255;not null" json:"status"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
