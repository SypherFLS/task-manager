package models

type User struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Email    string
	PasswordHASH string

	Tasks []Task `gorm:"foreignKey:UserID"`
}
