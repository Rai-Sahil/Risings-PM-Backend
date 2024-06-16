package models

import "time"

type Comment struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Content   string    `gorm:"type:text" json:"content"`
	UserID    int64     `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	TaskID    int64     `gorm:"not null" json:"task_id"`
	Task      Task      `gorm:"foreignKey:TaskID;references:ID"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
