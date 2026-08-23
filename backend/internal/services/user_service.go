package services

import (
	"context"
	"tm/internal/dto"
	_ "tm/internal/repository/models"
	"tm/internal/security"
	"tm/internal/validation"
)

func (s *Service) RegisterUser(ctx context.Context, user dto.RegisterDTO) error {
	if err := validation.Validate(user); err != nil {
		return err
	}

	hashed, err := security.HashPassword(user.Password)
	if err != nil {
		return err
	}

	userModel := dto.RegToModel(user, hashed)

	return s.repo.RegisterUser(ctx, userModel)
}

func (s *Service) LoginUser(ctx context.Context, user dto.LoginDTO) error {
	
	want, err := s.repo.LoginUser(ctx, user.Email)
	if err != nil {
		return err
	}

	return security.CheckPassword(user.Password, want)
}
