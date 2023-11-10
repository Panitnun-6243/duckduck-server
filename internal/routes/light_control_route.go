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

func LightControlRoutes(app *fiber.App) {
	app.Get("/api/v1/light-control", middlewares.Jwt(), getLightControlHandler)
	app.Put("/api/v1/light-control/:id", middlewares.Jwt(), updateLightControlHandler)
}

func getLightControlHandler(c *fiber.Ctx) error {
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))
	control, err := services.GetLightControlByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("Light control not found", err))
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(control))
}

func updateLightControlHandler(c *fiber.Ctx) error {
	var updatedControl models.LightControl
	if err := c.BodyParser(&updatedControl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	controlID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid light control ID", err))
	}

	// Extract userID from JWT claims
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	// Check if the user is authorized to update this light control
	lightControl, err := services.GetLightControlByIDAndUserID(controlID, userID)
	if err != nil || lightControl == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Unauthorized", nil))
	}

	err = services.UpdateUserLightControl(controlID, &updatedControl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Light control update failed", err))
	}

	return c.Status(fiber.StatusOK).JSON(responses.Info("Light control updated successfully"))
}
