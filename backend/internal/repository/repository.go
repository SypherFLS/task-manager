package repository

import (
	"context"
	"tm/internal/repository/models"
)

type Repository interface {
	Create(ctx context.Context,tasks []models.Task) error 
	GetAll(ctx context.Context) ([]models.Task, error)
	Update(ctx context.Context, id int, updates map[string]any) error 
	Delete(ctx context.Context, id int) error 
}