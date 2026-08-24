package dto

import "tm/internal/repository/models"

type RegisterDTO struct {
	Name     string `json:"name" validate:"required,min=5,max=15"`
	Email    string `json:"email" validate:"email,required,max=15,min=5"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"email,required,max=30"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

func RegToModel(user RegisterDTO, hash string) models.User {
	return models.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: hash,
	}
}
