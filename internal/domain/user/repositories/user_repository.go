package repositories

import (
	"context"

	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"

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

	err := repo.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *repositoryUser) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User

	err := repo.db.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *repositoryUser) Save(ctx context.Context, payload *model.User) (*model.User, error) {
	tx := repo.db.WithContext(ctx).Model(&payload)

	if err := tx.Save(&payload).Error; err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return payload, nil
}

func (repo *repositoryUser) Update(ctx context.Context, payload *model.User) (*model.User, error) {
	tx := repo.db.WithContext(ctx).Model(&payload)

	if err := tx.Save(&payload).Error; err != nil {
		logger.Error(err.Error())
		// custom error message
		return nil, err
	}

	return payload, nil
}

func (repo *repositoryUser) Delete(ctx context.Context, id string, user *model.User) error {
	tx := repo.db.WithContext(ctx).Model(&user)

	if err := tx.Delete(&user).Error; err != nil {
		logger.Error(err.Error())
		// custom error message
		return err
	}

	return nil
}
