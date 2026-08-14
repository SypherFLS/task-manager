package models

import (
	_ "time"
)

type Task struct {
	ID          int `gorm:"primaryKey"`
	Label       string
	Description string
	Priority    string
	// Duration time.Time
}
