package services

import (
	_ "fmt"
	"tm/internal/repository"
	_ "tm/internal/repository/models"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
