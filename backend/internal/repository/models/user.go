package models

type User struct {
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"unique"`
	Email        string `gorm:"unique"`
	PasswordHash string

	Tasks []Task `gorm:"foreignKey:UserID"`
}

type LoginResult struct {
	ID           int
	PasswordHash string
}
