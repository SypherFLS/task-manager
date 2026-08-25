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
}

type CreateTaskDTO struct {
	Label       string `json:"label" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=63"`
	Priority    Plevel `json:"priority" validate:"required,oneof=high middle low"`
}

type UpdateTaskDTO struct {
	Label       *string `json:"label" validate:"omitempty,min=3,max=100"`
	Description *string `json:"description" validate:"omitempty,max=63"`
	Priority    *Plevel `json:"priority" validate:"omitempty,oneof=high middle low"`
}

func ToModel(t CreateTaskDTO) *models.Task {
	return &models.Task{
		Label:       t.Label,
		Description: t.Description,
		Priority:    string(t.Priority),
	}
}

func ConvManyDto(t []CreateTaskDTO) []*models.Task {
	res := make([]*models.Task, 0, len(t))
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
		Priority:    Plevel(t.Priority),
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

func (t UpdateTaskDTO) ToMap() map[string]any {
	updates := make(map[string]any)

	if t.Label != nil {
		updates["label"] = *t.Label
	}

	if t.Description != nil {
		updates["description"] = *t.Description
	}

	if t.Priority != nil {
		updates["priority"] = string(*t.Priority)
	}

	return updates
}
