package models 

import (

)

type User struct {
	ID int `gorm:"primaryKey"`
	Name string
	Email string 
	Password string 

	Tasks []Task `gorm:"foreignKey:UserID"`
}