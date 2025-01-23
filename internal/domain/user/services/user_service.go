package services

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/hash"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/dtos"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
)

func (s *serviceUser) Finds(ctx context.Context) (*[]model.User, error) {
	res, err := s.userRepository.Finds(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		logger.Error(err.Error())
		return nil, err

	}

	return res, nil
}

func (s *serviceUser) FindByID(ctx context.Context, id string) (*model.User, error) {
	res, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		logger.Error(err.Error())
		return nil, err

	}

	return res, nil
}

func (s *serviceUser) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	res, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		logger.Error(err.Error())
		return nil, err
	}

	return res, nil
}

func (s *serviceUser) Save(ctx context.Context, payload *dto.CreateUserSchema) (*model.User, error) {
	now := time.Now()
	hashPassword, err := hash.HashPassword(payload.Password)
	if err != nil {
		// give a log error
		// Failed to hash password
		return nil, errors.New("error, can't process register")
	}

	// memasukkan payload ke dto/schema
	user := &model.User{
		ID:        uuid.New(),
		Name:      payload.Name,
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  hashPassword,
		Phone:     payload.Phone,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.userRepository.Save(ctx, user)
}

func (s *serviceUser) Update(ctx context.Context, payload *dto.UpdateUserSchema) (*model.User, error) {
	user, err := s.FindByID(ctx, payload.ID)
	if err != nil {
		return nil, err
	}
	user.Name = payload.Name
	user.Username = payload.Username
	user.Phone = payload.Phone
	user.IsActive = payload.IsActive

	return s.userRepository.Update(ctx, user)
}

func (s *serviceUser) Delete(ctx context.Context, id string) error {
	user, err := s.FindByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user id not found")
		}
		return err
	}

	return s.userRepository.Delete(ctx, id, user)
}
