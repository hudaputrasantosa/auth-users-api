package repositories

import (
	"context"

	"github.com/google/uuid"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"go.uber.org/zap"

	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
)

func (repo *repositoryUser) FindByUser(ctx context.Context, userId string) (*[]model.UsersActivityHistory, error) {
	var users *[]model.UsersActivityHistory
	id, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	err = repo.db.WithContext(ctx).Where("id = ?", id).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *repositoryUser) SaveActivity(ctx context.Context, payload *model.UsersActivityHistory) (*model.UsersActivityHistory, error) {
	tx := repo.db.WithContext(ctx).Model(&payload)

	if err := tx.Save(&payload).Error; err != nil {
		logger.Error("error insert", zap.Error(err))
		return nil, err
	}

	return payload, nil
}
