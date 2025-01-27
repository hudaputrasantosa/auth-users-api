package repositories

import (
	"context"

	"github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
)

func (repo *repositoryAuth) Register(ctx context.Context, payload models.User) (*models.User, error) {
	if err := repo.db.WithContext(ctx).Create(&payload).Error; err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &payload, nil
}
