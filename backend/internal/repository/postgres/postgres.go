package postgres

import (
	"context"
	"tm/internal/repository/models"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
)

type Repo struct {
	db *gorm.DB
}

var ErrNotFound = errors.New("record not found")

func InitDB() (*gorm.DB, error) {
	dsn := ""

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return db, err
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo {
		db : db,
	} 
}
 
func (r *Repo) GetAll(ctx context.Context) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	res := r.db.WithContext(ctx).Find(&tasks)
	if res.Error != nil {

		return nil, res.Error
	}

	return tasks, nil
}

func (r *Repo) Create(ctx context.Context,tasks []models.Task) error  {
	return r.db.WithContext(ctx).CreateInBatches(tasks, len(tasks)).Error
}
	
func (r *Repo) Update(ctx context.Context, id int, updates map[string]any) error {
	if len(updates)==0 {
		return nil
	}

	res := r.db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *Repo) Delete(ctx context.Context, id int) error {
	res := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Task{})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrNotFound
	}
	
	return nil
}