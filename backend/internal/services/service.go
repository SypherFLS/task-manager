package services

import (
	"tm/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo : repo,
	}
}

