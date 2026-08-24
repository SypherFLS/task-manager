package models

import (
	_ "time"
)

type Task struct {
	ID          int `gorm:"primaryKey"`
	Label       string
	Description string
	Priority    string

	UserID int
	User   User
	// Duration time.Time
}
