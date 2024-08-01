package models

type User struct {
	ID     int64  `gorm:"primary_key;auto_increment" json:"id"`
	Name   string `gorm:"size:255;not null" json:"name"`
	Email  string `gorm:"size:255;not null" json:"email"`
	Status string `gorm:"default:Active;not null" json:"status"`
}
