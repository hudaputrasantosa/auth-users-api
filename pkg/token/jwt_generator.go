package token

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils"
)

// Tokens struct to describe tokens object.
type Token struct {
	Token        string
	RefreshToken string
}

// GenerateNewTokens func for generate a new Access & Refresh tokens.
func GenerateNewToken(id string, minuteExpired int, tokenType utils.TokenTypeEnum) (*Token, error) {
	// Generate JWT Access token.
	accessToken, err := generateNewAccessToken(id, minuteExpired, tokenType)
	if err != nil {
		return nil, err
	}

	if tokenType == utils.AccessToken {
		// Generate JWT Refresh token.
		refreshToken, err := generateNewRefreshToken()
		if err != nil {
			return nil, err
		}

		return &Token{
			Token:        accessToken,
			RefreshToken: refreshToken,
		}, nil
	}

	return &Token{
		Token: accessToken,
	}, nil
}

// private func for generate a new Access token.
func generateNewAccessToken(userID string, minuteExpired int, tokenType utils.TokenTypeEnum) (string, error) {
	// Create a new claims to JWT.
	claims := jwt.MapClaims{
		"id":   userID,
		"type": tokenType,
		"exp":  time.Now().Add(time.Minute * time.Duration(minuteExpired)).Unix(),
	}

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	tokenResult, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenResult, nil
}

// private func for generate a new refresh access token.
func generateNewRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	hash := sha256.New()

	// Create a new now date and time string with salt.
	refresh := config.Config("JWT_REFRESH_KEY") + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := hash.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	// Set expires hours count for refresh key from .env file.
	hoursCount, _ := strconv.Atoi(config.Config("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT"))

	// Set expiration time.
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}
