package postgres

import (
	"context"
	"tm/internal/repository/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
)

type Repo struct {
	db *gorm.DB
}

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
	var tasks []models.Task

	res := r.db.Find(&tasks)
	if res.Error != nil {
		return nil, res.Error
	}

	return tasks, nil
}

func (r *Repo) Create(ctx context.Context,task models.Task) error  {
	
	return nil
}
	
func (r *Repo) Update(ctx context.Context, task models.Task) error {
	return nil
}
func (r *Repo) Delete(ctx context.Context, id int) error {
	return nil
}