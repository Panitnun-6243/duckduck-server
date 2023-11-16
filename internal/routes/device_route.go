package routes

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/responses"
	"github.com/Panitnun-6243/duckduck-server/internal/services"
	"github.com/gofiber/fiber/v2"
)

func DeviceRoutes(app *fiber.App) {
	app.Post("/api/v1/device-login", deviceLoginHandler)
}

func deviceLoginHandler(c *fiber.Ctx) error {
	var deviceLoginRequest models.DeviceLoginRequest
	if err := c.BodyParser(&deviceLoginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	token, err := services.DeviceLogin(deviceLoginRequest.DeviceCode, deviceLoginRequest.Secret)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Device login failed", err))
	}

	return c.Status(fiber.StatusOK).JSON(responses.Info(map[string]string{
		"token": token,
	}))
}
