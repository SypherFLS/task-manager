package services

import (
	"tm/internal/dto"
	"tm/internal/api/utils/params"
	_ "fmt"
	_ "tm/internal/repository/models"
	"context"
)

var holder = 0 // заменить на айдишник


func (s *Service) CreateTask(ctx context.Context, tasksDto []dto.CreateTaskDTO) error {

	tasks := dto.ConvManyDto(tasksDto)

	return s.repo.CreateTask(ctx, tasks, holder)
}

func (s *Service) ReadTask(ctx context.Context, query params.FilPage) ([]dto.TaskDTO, error)  {
	data, err := s.repo.GetTask(ctx, query, holder)
	req := dto.ToDTOs(data) 

	return req, err
}

func (s *Service) DeleteTask(ctx context.Context, id int) error{
	return s.repo.DeleteTask(ctx, id, holder)
}	

func (s *Service) UpdateTask(ctx context.Context,id int, updto dto.UpdateTaskDTO) error { 
	updates := updto.ToMap()

	if len(updates) == 0 {
		return nil // маппинг ошибок
	}

	return s.repo.UpdateTask(ctx, id, updates, holder)
}