package postgres

import (
	"tm/internal/repository/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repo struct {
	db *gorm.DB
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Task{}, &models.User{}); err != nil {
		return nil, err
	}

	return db, err
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}
