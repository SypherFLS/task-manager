package dto

import (
	"time"
	"tm/internal/repository/models"
)

type TaskDTO struct {
	ID          int       `json:"id"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	Duration    time.Time `json:"duration"`
}

type CTaskDTO struct {
	Label       string    `json:"label" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"max=63"`
	Duration    time.Time `json:"duration"`
}

func ToModel(t CTaskDTO) models.Task {
	return models.Task{
		Label:       t.Label,
		Description: t.Description,
		Duration:    t.Duration,
	}
}

func ConvManyDto(t []CTaskDTO) []models.Task {
	res := make([]models.Task, 0, len(t))
	for _, r := range t {
		res = append(res, ToModel(r))
	}

	return res
}

func ToDTO(t models.Task) TaskDTO {
	return TaskDTO{
		Label:       t.Label,
		Description: t.Description,
		Duration:    t.Duration,
	}
}

func ToDTOs(tasks []models.Task) []TaskDTO {
	res := make([]TaskDTO, 0, len(tasks))

	for _, task := range tasks {
		res = append(res, ToDTO(task))
	}

	return res
}
