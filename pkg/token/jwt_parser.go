package token

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
)

// Token Metadata struct to describe token Metadata object.
type TokenMetadata struct {
	UserID      uuid.UUID
	Expires     int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}

		// Expires time.
		expires := int64(claims["exp"].(float64))

		return &TokenMetadata{
			UserID:      userID,
			Expires:     expires,
		}, nil
	}

	return nil, err
}

// private func to verify token
func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, JwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// private func to extract token from Header
func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func JwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.Config("JWT_SECRET")), nil
}

// ParseRefreshToken func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}

// another version to verify token jwt
// func VerifyToken(tokenString string) (bool, error) {
//  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//   return []byte(os.Getenv("JWT_SECRET")), nil
//  })
//  if err != nil {
//   return false, err
//  }

//  return token.Valid, nil
// }
