package models

import (
	_ "time"
)

type Task struct {
	ID          int `gorm:"primaryKey"`
	Label       string 
	Description string
	Priority    string

	User User `gorm:"foreignKey:UserID;references:ID"`
	// Duration time.Time
}
