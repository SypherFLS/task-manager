package postgres

import (
	"context"
	"tm/internal/repository/models"
)

func (r *Repo) RegisterUser(ctx context.Context, user models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
