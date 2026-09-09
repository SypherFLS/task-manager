package models

import (
	_ "time"
)

type Task struct {
	ID          int    `gorm:"primaryKey"`
	Label       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Priority    string `gorm:"oneof:high,mid,low"`

	UserID int
	User   User
}
