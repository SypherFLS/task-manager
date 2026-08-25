package postgres

import (
	"context"
	"tm/internal/repository/models"
)

func (r *Repo) RegisterUser(ctx context.Context, user models.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (r *Repo) LoginUser(ctx context.Context, email string) (*models.LoginResult, error) {
	var result models.LoginResult
	res := r.db.WithContext(ctx).Model(&models.User{}).Select("id", "password_hash").Where("email = ?", email).First(&result)
	return &result, res.Error
}
