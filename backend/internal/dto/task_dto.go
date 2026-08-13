package dto

import (
	_ "time"
	"tm/internal/repository/models"
)

type Plevel string

const (
	High   Plevel = "high"
	Middle Plevel = "middle"
	Low    Plevel = "low"
)

type TaskDTO struct {
	ID          int    `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Priority    Plevel `json:"priority"`
	// Duration    time.Time `json:"duration"` // TODO time.Format validation
}

type CreateTaskDTO struct {
	Label       string `json:"label" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=63"`
	Priority    Plevel `json:"priority" validate:"oneof=high middle low"`
	// Duration    time.Time `json:"duration"`
}

type UpdateTaskDTO struct {
	Label       string `json:"label" validate:"min=3,max=100"`
	Description string `json:"description" validate:"max=63"`
	Priority    Plevel `json:"priority" validate:"oneof=high middle low"`
	// Duration    time.Time `json:"duration"`
}

func ToModel(t CreateTaskDTO) models.Task {
	return models.Task{
		Label:       t.Label,
		Description: t.Description,
		// Priority:    t.Priority,
		// Duration:    t.Duration,
	}
}

func ConvManyDto(t []CreateTaskDTO) []models.Task {
	res := make([]models.Task, 0, len(t))
	for _, r := range t {
		res = append(res, ToModel(r))
	}

	return res
}

func ToDTO(t models.Task) TaskDTO {
	return TaskDTO{
		ID:          t.ID,
		Label:       t.Label,
		Description: t.Description,
		// Duration:    t.Duration,
	}
}

func ToDTOs(tasks []models.Task) []TaskDTO {
	res := make([]TaskDTO, 0, len(tasks))

	for _, task := range tasks {
		res = append(res, ToDTO(task))
	}

	return res
}
