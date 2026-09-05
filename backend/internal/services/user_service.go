package services

import (
	"context"
	"tm/internal/auth"
	"tm/internal/dto"
	"tm/internal/validation"
)

func (s *Service) RegisterUser(ctx context.Context, user dto.RegisterDTO) error {
	if err := validation.Validate(user); err != nil {
		return err
	}

	hashed, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	userModel := dto.RegToModel(user, hashed)

	return s.repo.RegisterUser(ctx, userModel)
}

func (s *Service) LoginUser(ctx context.Context, user dto.LoginDTO) (string, error) {
	result, err := s.repo.LoginUser(ctx, user.Email)

	if err != nil {
		return "", err
	}

	if err := auth.CheckPassword(user.Password, result.PasswordHash); err != nil {
		return "", err
	}

	token, err := s.jwtManager.Generate(result.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
