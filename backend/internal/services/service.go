package services

import (
	"tm/internal/repository"
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

func (s *Service) Create(ctx context.Context, ) error {


	return nil
}