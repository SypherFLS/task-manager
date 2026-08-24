package services

import (
	"context"
	"tm/internal/api/utils/params"
	"tm/internal/dto"
	_ "tm/internal/repository/models"
)

func (s *Service) CreateTask(ctx context.Context, tasksDto []dto.CreateTaskDTO, userID int) error {

	tasks := dto.ConvManyDto(tasksDto)

	return s.repo.CreateTask(ctx, tasks, userID)
}

func (s *Service) ReadTask(ctx context.Context, query params.FilPage, userID int) ([]dto.TaskDTO, error) {
	data, err := s.repo.GetTask(ctx, query, userID)
	req := dto.ToDTOs(data)

	return req, err
}

func (s *Service) DeleteTask(ctx context.Context, taskID int, userID int) error {
	return s.repo.DeleteTask(ctx, taskID, userID)
}

func (s *Service) UpdateTask(ctx context.Context, taskID int, updto dto.UpdateTaskDTO, userID int) error {
	updates := updto.ToMap()

	if len(updates) == 0 {
		return nil // маппинг ошибок
	}

	return s.repo.UpdateTask(ctx, taskID, updates, userID)
}
