package models

type Student struct {
	ID         int64  `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `gorm:"size:255;not null" json:"name"`
	AssigneeID int64  `gorm:"not null" json:"assignee_id"`
	Assignee   User   `gorm:"foreignKey:AssigneeID;references:ID"`
}
