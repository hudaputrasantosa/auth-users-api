package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
)

// Function untuk mendeteksi jenis perangkat berdasarkan User-Agent
func DetectDevice(ctx *fiber.Ctx) models.UsersActivityHistory {
	device := "web-desktop"
	userAgent := strings.ToLower(ctx.Get("User-Agent"))

	if strings.Contains(userAgent, "mobile") {
		device = "web-mobile"
	}

	if strings.Contains(userAgent, "postman") || strings.Contains(userAgent, "curl") || strings.Contains(userAgent, "insomnia") {
		device = "api"
	}

	return models.UsersActivityHistory{
		Context:     models.Login,
		Ip:          ctx.IP(),
		AgentClient: userAgent,
		Device:      device,
	}
}
