package routes

import (
	"encoding/json"
	"fmt"
	"github.com/Panitnun-6243/duckduck-server/internal/middlewares"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/responses"
	"github.com/Panitnun-6243/duckduck-server/internal/services"
	"github.com/Panitnun-6243/duckduck-server/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HslLightRoutes(app *fiber.App) {
	app.Get("/api/v1/hsl-light", middlewares.Jwt(), getHslLightHandler)
	app.Patch("/api/v1/hsl-light/:id", middlewares.Jwt(), updateHslLightHandler)
}

func getHslLightHandler(c *fiber.Ctx) error {
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	hslLight, err := services.GetHslLightByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("Cannot get hsl color", err))
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(bson.M{"hsl_color": hslLight}))
}

func updateHslLightHandler(c *fiber.Ctx) error {
	var requestBody struct {
		HslColor models.Hsl `json:"hsl_color"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	hslID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid hsl light ID", err))
	}

	// Extract userID from JWT claims
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	// Check if the user is authorized to update this light control
	_, err = services.GetHslLightByIDAndUserID(hslID, userID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Unauthorized", nil))
	}

	err = services.UpdateUserHslLight(hslID, requestBody.HslColor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Hsl light update failed", err))
	}

	// Publish the update to MQTT
	deviceCode, err := services.GetDeviceCodeByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Failed to get device code", err))
	}
	mqttTopic := fmt.Sprintf("%s/hsl", deviceCode)
	filter := bson.M{
		"h": requestBody.HslColor.Hue,
		"s": requestBody.HslColor.Saturation,
		"l": requestBody.HslColor.Lightness,
	}
	payload, _ := json.Marshal(filter) // Convert the updatedControl struct to JSON
	client := util.CreateMqttClient()
	util.Publish(client, mqttTopic, string(payload))

	return c.Status(fiber.StatusOK).JSON(responses.Info("Hsl light updated successfully"))
}
