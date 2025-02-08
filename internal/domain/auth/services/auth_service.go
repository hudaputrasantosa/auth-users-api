package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/utils"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"github.com/hudaputrasantosa/auth-users-api/pkg/hash"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	"github.com/hudaputrasantosa/auth-users-api/pkg/notification"
	"github.com/hudaputrasantosa/auth-users-api/pkg/otp"
	"github.com/hudaputrasantosa/auth-users-api/pkg/token"
	globalUtils "github.com/hudaputrasantosa/auth-users-api/pkg/utils"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/cache"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/templates"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserTokenResponse struct {
	AccessToken  string
	RefreshToken string
	Expired      int
}

type UserRegisterResponse struct {
	Email   string
	Token   string
	Expired int
}

func (s *serviceAuth) ValidateUser(ctx context.Context, payload dto.ValidateUserSchema) (*UserTokenResponse, int, error) {
	// Create user model to store data
	var user *model.User

	// check email exist
	user, err := s.userRepository.FindByEmail(ctx, payload.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error("email not found or not registered")
			return nil, fiber.StatusNotFound, utils.FailedUserLogin
		}

		logger.Error(err.Error())
		return nil, fiber.StatusInternalServerError, err
	}

	//chack password
	isPassword := hash.CheckPasswordHash(payload.Password, user.Password)
	if !isPassword {
		logger.Error("password is not match")
		return nil, fiber.StatusConflict, utils.FailedUserLogin
	}

	// check status user
	if !user.IsActive {
		// send email service to verification
		otp := otp.GenerateOTP(6)

		key := fmt.Sprintf(cache.REGISTER_OTP+"%s", user.ID.String()+"_"+otp)
		s.redis.Set(ctx, key, otp, 5*time.Minute)

		minuteExpired := 5
		token, err := token.GenerateNewToken(user.ID.String(), minuteExpired, globalUtils.VerificationToken)
		if err != nil {
			logger.Error("Error login", zap.Error(err))
			return nil, fiber.StatusInternalServerError, utils.FailedUserLogin
		}

		_, err = notification.MailersendNotification(&notification.RecipientInformation{
			Email: user.Email,
			Name:  user.Name,
		}, &templates.DataBodyInformation{
			Name:            user.Name,
			Otp:             otp,
			MessageTemplate: templates.Otp_template,
		})
		if err != nil {
			logger.Error("Failed send notification")
		}

		res := &UserTokenResponse{
			AccessToken: token.Token,
			Expired:     minuteExpired,
		}
		logger.Error("User not active, Please to verification. check your email")
		return res, fiber.StatusBadRequest, errors.New("please to verification email")
	}

	// generate new access token and refresh token jwt
	minuteExpired, _ := strconv.Atoi(config.Config("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	userToken, err := token.GenerateNewToken(user.ID.String(), minuteExpired, globalUtils.AccessToken)
	if err != nil {
		logger.Error("failed generate token", zap.Error(err))
		return nil, fiber.StatusInternalServerError, globalUtils.ErrorGlobalPublicMessage
	}

	res := &UserTokenResponse{
		AccessToken:  userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}

	return res, fiber.StatusOK, err
}

func (s *serviceAuth) RegisterUser(ctx context.Context, payload dto.RegisterUserSchema) (interface{}, int, error) {
	// check email existing
	user, err := s.userRepository.FindByEmail(ctx, payload.Email)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.Error("Error register", zap.Error(err))
			return nil, fiber.StatusInternalServerError, utils.ErrorUserRegister
		}
	} else if user != nil {
		return nil, fiber.StatusConflict, utils.FailedUserRegister
	}

	// if admin role, get header secret key for create admin user
	now := time.Now()

	hashPassword, err := hash.HashPassword(payload.Password)
	if err != nil {
		logger.Error("Error, failed to hash password", zap.Error(err))
		return nil, fiber.StatusInternalServerError, utils.ErrorUserRegister
	}

	createUser := model.User{
		ID:        uuid.New(),
		Name:      payload.Name,
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  hashPassword,
		Phone:     payload.Phone,
		CreatedAt: now,
		UpdatedAt: now,
	}

	user, err = s.authRepository.Register(ctx, createUser)
	if err != nil {
		logger.Error("Error register", zap.Error(err))
		return nil, fiber.StatusInternalServerError, utils.ErrorUserRegister
	}
	// generate otp
	otp := otp.GenerateOTP(6)

	// set otp code use redis
	key := fmt.Sprintf(cache.REGISTER_OTP+"%s", user.ID.String()+"_"+otp)
	s.redis.Set(ctx, key, otp, 5*time.Minute)

	// generate otp token from jwt
	minuteExpired := 5
	verifyToken, err := token.GenerateNewToken(user.ID.String(), minuteExpired, globalUtils.VerificationToken)
	if err != nil {
		logger.Error("Error register", zap.Error(err))
		return nil, fiber.StatusInternalServerError, utils.ErrorUserRegister
	}

	// sent otp to active email that registered
	_, err = notification.MailersendNotification(&notification.RecipientInformation{
		Email: user.Email,
		Name:  user.Name,
	}, &templates.DataBodyInformation{
		Name:            user.Name,
		Otp:             otp,
		MessageTemplate: templates.Otp_template,
	})
	if err != nil {
		logger.Error("Failed send notification")
	}

	res := &UserRegisterResponse{
		Email:   user.Email,
		Token:   verifyToken.Token,
		Expired: minuteExpired,
	}

	// return success with token otp
	return res, fiber.StatusCreated, nil
}

func (s *serviceAuth) VerificationUser(ctx context.Context, payload dto.VerificationUser) (string, int, error) {
	// parse for valid token
	token, err := jwt.Parse(payload.Token, token.JwtKeyFunc)
	if err != nil {
		logger.Error("error jwt", zap.Error(err))
		return "", fiber.StatusBadRequest, utils.ErrorUserVerification
	}
	// get userid by claims token
	claims, ok := token.Claims.(jwt.MapClaims)["id"]
	if !ok {
		return "", fiber.StatusBadRequest, utils.ErrorUserVerification
	}

	// make a key otp verify for get redis by key
	userId := fmt.Sprintf("%v", claims)
	key := fmt.Sprintf(cache.REGISTER_OTP+"%s", userId+"_"+payload.Otp)
	userOtp := s.redis.Get(ctx, key).Val()

	// check for valid otp
	if userOtp == "" || userOtp != payload.Otp {
		return "", fiber.StatusBadRequest, errors.New("can't process verification, because otp not valid or other")
	}

	res, err := s.userRepository.UpdateStatusById(ctx, userId)
	if err != nil {
		logger.Error("error update status user", zap.Error(err))
		return "", fiber.StatusInternalServerError, utils.ErrorUserVerification
	}

	// delete key otp not used
	s.redis.Del(ctx, key)

	return res.Email, fiber.StatusOK, nil
}
