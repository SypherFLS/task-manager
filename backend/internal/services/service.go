package services

import (
	"tm/internal/auth"
	"tm/internal/repository"
)

type Service struct {
	repo       repository.Repository
	jwtManager *auth.JWTManager
}

func NewService(repo repository.Repository, JWTManager *auth.JWTManager) *Service {
	return &Service{
		repo:       repo,
		jwtManager: JWTManager,
	}
}
