package repository

import (
	"context"
	"tm/internal/repository/models"
)

type Repository interface {
	Create(ctx context.Context,task models.Task) error 
	GetAll(ctx context.Context) ([]models.Task, error)
	Update(ctx context.Context, task models.Task) error 
	Delete(ctx context.Context, id int) error 
}