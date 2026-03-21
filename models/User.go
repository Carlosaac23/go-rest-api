package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID        string
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Tasks     []Task
}
