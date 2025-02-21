package services

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hudaputrasantosa/auth-users-api/pkg/hash"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/dtos"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/user/utils"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
)

func (s *serviceUser) Finds(ctx context.Context) (*[]model.User, int, error) {
	res, err := s.userRepository.Finds(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, fiber.StatusNotFound, nil
		}

		logger.Error(err.Error())
		return nil, fiber.StatusInternalServerError, err
	}

	return res, fiber.StatusOK, nil
}

func (s *serviceUser) FindByID(ctx context.Context, id string) (*model.User, int, error) {
	res, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.StatusNotFound, errors.New("user not found")
		}
		logger.Error(err.Error())
		return nil, fiber.StatusInternalServerError, err

	}

	return res, fiber.StatusOK, nil
}

func (s *serviceUser) FindByEmail(ctx context.Context, email string) (*model.User, int, error) {
	res, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.StatusNotFound, errors.New("user not found")
		}
		logger.Error(err.Error())
		return nil, fiber.StatusInternalServerError, err
	}

	return res, fiber.StatusOK, nil
}

func (s *serviceUser) Save(ctx context.Context, payload *dto.CreateUserSchema) (*model.User, int, error) {
	now := time.Now()

	hashPassword, err := hash.HashPassword(payload.Password)
	if err != nil {
		logger.Error("Error, failed to hash password", zap.Error(err))
		return nil, fiber.StatusInternalServerError, utils.ErrorUserCreated
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

	data, err := s.userRepository.Save(ctx, user)
	if err != nil {
		logger.Error("Error, failed to save data user", zap.Error(err))
		return nil, fiber.StatusInternalServerError, utils.ErrorUserCreated
	}

	return data, fiber.StatusCreated, nil
}

func (s *serviceUser) Update(ctx context.Context, id string, payload *dto.UpdateUserSchema) (*model.User, int, error) {
	user, status, err := s.FindByID(ctx, payload.ID)
	if err != nil {
		return nil, status, err
	}
	user.Name = payload.Name
	user.Username = payload.Username
	user.Phone = payload.Phone
	user.IsActive = payload.IsActive

	data, err := s.userRepository.Update(ctx, user)
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.New("failed update user")
	}

	return data, fiber.StatusOK, nil
}

func (s *serviceUser) Delete(ctx context.Context, id string) (int, error) {
	user, status, err := s.FindByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return status, errors.New("user id not found")
		}
		return status, err
	}

	err = s.userRepository.Delete(ctx, id, user)
	if err != nil {
		return fiber.StatusInternalServerError, errors.New("failed delete user")
	}

	return fiber.StatusOK, nil
}
