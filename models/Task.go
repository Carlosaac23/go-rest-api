package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	ID          string
	Title       string `gorm:"type:varchar(100);not null;unique"`
	Description string
	Done        bool   `gorm:"default:false"`
	UserID      string `json:"user_id"`
}
