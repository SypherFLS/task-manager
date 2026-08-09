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

func (s *Service) Create(ctx context.Context, tasksDto []dto.TaskDTO) error {

	tasks := dto.ConvManyDto(tasksDto)

	err := s.repo.Create(ctx, tasks)

	return err
}

func (s *Service) Read(ctx context.Context) (error)  {

	return nil
}

func (s *Service) Delete(ctx context.Context, id int) error{
	return s.repo.Delete(ctx, id)
}	

func (s *Service) Update(ctx context.Context) error {

	return nil
}