package services

import (
	"tm/internal/dto"
	"tm/internal/repository"
	_ "tm/internal/repository/models"
	"context"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo : repo,
	}
}

func (s *Service) Create(ctx context.Context, tasksDto []dto.CTaskDTO) error {

	tasks := dto.ConvManyDto(tasksDto)

	return s.repo.Create(ctx, tasks)
}

func (s *Service) Read(ctx context.Context) ([]dto.TaskDTO, error)  {
	data, err := s.repo.GetAll(ctx)
	req := dto.ToDTOs(data) 

	return req, err
}

func (s *Service) Delete(ctx context.Context, id int) error{
	return s.repo.Delete(ctx, id)
}	

func (s *Service) Update(ctx context.Context,id int, cdto dto.CTaskDTO) error {

	return nil
}