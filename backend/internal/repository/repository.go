package repository

import (
	"context"
	"tm/internal/repository/models"
	"tm/internal/api/utils/params"
)

type Repository interface {
	CreateTask(ctx context.Context,tasks []models.Task, userID int) error 
	GetTask(ctx context.Context,query params.FilPage, userID int) ([]models.Task, error)
	UpdateTask(ctx context.Context, id int, updates map[string]any, userID int) error 
	DeleteTask(ctx context.Context, id int, userID int) error 
}