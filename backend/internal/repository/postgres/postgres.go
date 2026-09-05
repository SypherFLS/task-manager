package postgres

import (
	"fmt"
	"tm/internal/config"
	"tm/internal/repository/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repo struct {
	db *gorm.DB
}

func InitDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
