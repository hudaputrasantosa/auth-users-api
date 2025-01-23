package repositories

import (
	"context"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
)

func (repo *repositoryAuth) Register(ctx context.Context, payload *dto.RegisterUserSchema) (interface{}, error) {
	tx := repo.db.WithContext(ctx).Model(&payload)

	if err := tx.Save(&payload).Error; err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return payload, nil
}
