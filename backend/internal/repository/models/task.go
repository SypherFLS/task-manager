package models

import (
	"time"
)

type Task struct {
	ID int `gorm:"primaryKey"`
	Label string
	Description string
	Duration time.Time
}