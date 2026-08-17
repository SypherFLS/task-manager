package repository

import (
	"context"
	"tm/internal/repository/models"
	"tm/internal/api/utils/params"
)

type Repository interface {
	Create(ctx context.Context,tasks []models.Task) error 
	Get(ctx context.Context,query params.FilPage) ([]models.Task, error)
	Update(ctx context.Context, id int, updates map[string]any) error 
	Delete(ctx context.Context, id int) error 
}