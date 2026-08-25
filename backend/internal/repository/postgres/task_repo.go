package postgres

import (
	"context"
	"tm/internal/api/utils/params"
	"tm/internal/apperrors"
	"tm/internal/repository/models"
)

func (r *Repo) GetTask(ctx context.Context, query params.FilPage, userID int) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	res := r.db.WithContext(ctx).Model(&models.Task{}).Where("user_id = ?", userID)

	if query.Priority != nil {
		res = res.Where("priority = ?", *query.Priority)
	}

	res = res.Limit(query.Limit).Offset(query.Offset)

	res = res.Find(&tasks)

	if res.Error != nil {
		return nil, res.Error
	}

	return tasks, nil
}

func (r *Repo) CreateTask(ctx context.Context, tasks []*models.Task, userID int) error {
	for _, task := range tasks {
		task.UserID = userID
	}
	return r.db.WithContext(ctx).CreateInBatches(tasks, len(tasks)).Error
}

func (r *Repo) UpdateTask(ctx context.Context, taskID int, updates map[string]any, userID int) error {
	if len(updates) == 0 {
		return nil
	}

	res := r.db.WithContext(ctx).Model(&models.Task{}).Where("id = ? AND user_id = ?", taskID, userID).Updates(updates)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}

func (r *Repo) DeleteTask(ctx context.Context, taskID int, userID int) error {
	res := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}
