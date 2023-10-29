package routes

import (
	"github.com/Panitnun-6243/duckduck-server/internal/middlewares"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/responses"
	"github.com/Panitnun-6243/duckduck-server/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DashboardRoutes(app *fiber.App) {
	app.Get("/api/v1/dashboard-config", middlewares.Jwt(), getDashboardConfigHandler)
	app.Put("/api/v1/dashboard-config/:id", middlewares.Jwt(), updateDashboardConfigHandler)
}

func getDashboardConfigHandler(c *fiber.Ctx) error {
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	config, err := services.GetDashboardConfigByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("Dashboard config not found", err))
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(config))
}

func updateDashboardConfigHandler(c *fiber.Ctx) error {
	var updatedConfig models.DashboardConfig
	if err := c.BodyParser(&updatedConfig); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	configID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid dashboard config ID", err))
	}

	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	config, err := services.GetDashboardConfigByUser(userID)
	if err != nil || config == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Unauthorized", nil))
	}

	err = services.UpdateUserDashboardConfig(configID, &updatedConfig)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Dashboard config update failed", err))
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info("Dashboard config updated successfully"))
}
