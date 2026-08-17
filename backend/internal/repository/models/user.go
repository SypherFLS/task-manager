package models 

import (

)

type User struct {
	ID int `gorm:"primaryKey"`
	Name string
	Contact string 
	Password string 
}