package models

type User struct {
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"unique"`
	Email        string `gorm:"unique"`
	PasswordHASH string

	Tasks []Task `gorm:"foreignKey:UserID"`
}
