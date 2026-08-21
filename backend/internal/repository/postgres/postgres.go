package postgres

import (
	"context"
	"tm/internal/repository/models"
	"tm/internal/apperrors"
	"tm/internal/api/utils/params"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
)

type Repo struct {
	db *gorm.DB
}


func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{ // поменять реализацию на конфиг
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
	return &Repo {
		db : db,
	} 
}
 
func (r *Repo) GetTask(ctx context.Context, query params.FilPage, userID int) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	res := r.db.WithContext(ctx).Model(&models.Task{})

	if query.Priority != nil {
		res = res.Where("priority = ?", *query.Priority)
	}

	res = res.Limit(query.Limit).Offset(query.Offset)

	res = res.Find(&tasks)

	if res.Error != nil {
		return nil, res.Error
	}

	return tasks, nil
}



func (r *Repo) CreateTask(ctx context.Context,tasks []models.Task, userID int) error  {
	return r.db.WithContext(ctx).CreateInBatches(tasks, len(tasks)).Error
}
	
func (r *Repo) UpdateTask(ctx context.Context, id int, updates map[string]any, userID int) error {
	if len(updates)==0 {
		return nil
	}

	res := r.db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", id).Updates(updates)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}

func (r *Repo) DeleteTask(ctx context.Context, id int, userID int) error {
	res := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Task{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	
	return nil
}