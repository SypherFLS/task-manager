package services

import (
	"tm/internal/repository"
	_ "fmt"
	_ "tm/internal/repository/models"
)


type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo : repo,
	}
}

