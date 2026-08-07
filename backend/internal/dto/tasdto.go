package dto

import (
	"time"
	"tm/internal/repository/models"
)

type TaskDTO struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Duration    time.Time `json:"duration"`
}

func ToModel(t TaskDTO) models.Task {
	return models.Task {
		Label : t.Label,
		Description: t.Description,
		Duration : t.Duration,
	}
}

func ConvManyDto(t []TaskDTO) []models.Task {
	res := make([]models.Task, 0, len(t))
	for _, r := range t {
		res = append(res, ToModel(r))
	}

	return res
}