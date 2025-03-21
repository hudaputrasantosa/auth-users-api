package services

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/dtos"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"github.com/hudaputrasantosa/auth-users-api/pkg/hash"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils"
	"go.uber.org/zap"
)

func (s *serviceUser) UpdateMemberById(ctx context.Context, id string, payload *dto.UpdateUserMemberSchema) (*model.User, int, error) {
	user, status, err := s.FindByID(ctx, id)
	if err != nil {
		return nil, status, err
	}
	user.Name = payload.Name
	user.Phone = payload.Phone

	data, err := s.userRepository.Update(ctx, user)
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.New("failed update profile")
	}

	return data, fiber.StatusOK, nil
}

func (s *serviceUser) UpdatePassword(ctx context.Context, id string, payload *dto.UpdatePasswordMember) (int, error) {
	user, status, err := s.FindByID(ctx, id)
	if err != nil {
		return status, err
	}

	//check last password
	if !hash.CheckPasswordHash(payload.LastPassword, user.Password) {
		return fiber.StatusConflict, errors.New("password is not match")
	}

	// update new password by userId
	newHashPassword, err := hash.HashPassword(payload.NewPassword)
	if err != nil {
		logger.Error("Error, failed to hash password", zap.Error(err))
		return fiber.StatusInternalServerError, utils.FailedUpdatePassword
	}

	user.Password = newHashPassword
	user.LastPasswordUpdated = time.Now()

	_, err = s.userRepository.Update(ctx, user)
	if err != nil {
		return fiber.StatusInternalServerError, utils.FailedUpdatePassword
	}

	return fiber.StatusOK, nil
}

func (s *serviceUser) DeactivatedAccount(ctx context.Context, id string, password string) (int, error) {
	user, status, err := s.FindByID(ctx, id)
	if err != nil {
		return status, err
	}

	//chack password
	if !hash.CheckPasswordHash(password, user.Password) {
		logger.Error("password is not match")
		return fiber.StatusConflict, errors.New("password is not match")
	}

	err = s.userRepository.Delete(ctx, *user)
	if err != nil {
		return fiber.StatusInternalServerError, errors.New("failed deactivated account")
	}

	return fiber.StatusOK, nil
}
