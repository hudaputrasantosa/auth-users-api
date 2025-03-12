package repositories

import (
	"context"

	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"go.uber.org/zap"

	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
)

func (repo *repositoryUser) Finds(ctx context.Context) (*[]model.User, error) {
	var users *[]model.User

	err := repo.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *repositoryUser) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user *model.User

	err := repo.db.WithContext(ctx).Select("id", "name", "username", "email", "phone", "is_active").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *repositoryUser) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User

	err := repo.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *repositoryUser) Save(ctx context.Context, payload *model.User) (*model.User, error) {
	tx := repo.db.WithContext(ctx).Model(&payload)

	if err := tx.Save(&payload).Error; err != nil {
		logger.Error("error insert", zap.Error(err))
		return nil, err
	}

	return payload, nil
}

func (repo *repositoryUser) Update(ctx context.Context, payload *model.User) (*model.User, error) {
	tx := repo.db.WithContext(ctx).Model(&payload)

	if err := tx.Save(&payload).Error; err != nil {
		logger.Error("error update", zap.Error(err))
		return nil, err
	}

	return payload, nil
}

func (repo *repositoryUser) UpdateStatusById(ctx context.Context, userId string) (*model.User, error) {
	user, err := repo.FindByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	user.IsActive = true
	tx := repo.db.WithContext(ctx).Model(&user)

	if err := tx.Save(&user).Error; err != nil {
		logger.Error("error update", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (repo *repositoryUser) Delete(ctx context.Context, user model.User) error {
	tx := repo.db.WithContext(ctx).Model(&user)

	if err := tx.Delete(&user).Error; err != nil {
		logger.Error("error delete", zap.Error(err))
		return err
	}

	return nil
}
