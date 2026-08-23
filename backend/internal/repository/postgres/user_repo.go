package postgres

import (
	"context"
	"tm/internal/repository/models"
)

func (r *Repo) RegisterUser(ctx context.Context, user models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}


func (r *Repo) LoginUser(ctx context.Context, email string) (string, error){
	var password string
	res := r.db.WithContext(ctx).Where("email = ?", email).Find(&password)
	return password, res.Error
}