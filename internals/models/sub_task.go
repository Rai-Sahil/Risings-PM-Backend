package models

type SubTask struct {
	ID         int64  `gorm:"primary_key;auto_increment" json:"id"`
	Title      string `gorm:"size:255;not null" json:"title"`
	AssigneeID int64  `gorm:"not null" json:"assignee_id"`
	Assignee   User   `gorm:"foreignKey:AssigneeID;references:ID"`
	TaskID     int64  `gorm:"not null" json:"task_id"`
	Status     string `gorm:"size:255;default:'Pending'" json:"status"`
	Task       Task   `gorm:"foreignKey:TaskID;references:ID"`
}
