package models

import (
	_ "time"
)

type Plevel string

const (
	High   Plevel = "high"
	Middle Plevel = "middle"
	Low    Plevel = "low"
)

type Task struct {
	ID          int `gorm:"primaryKey"`
	Label       string
	Description string
	Priority    Plevel
	// Duration time.Time
}
