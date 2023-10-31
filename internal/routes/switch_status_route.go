package routes

import (
	"encoding/json"
	"fmt"
	"github.com/Panitnun-6243/duckduck-server/internal/middlewares"
	"github.com/Panitnun-6243/duckduck-server/internal/responses"
	"github.com/Panitnun-6243/duckduck-server/internal/services"
	"github.com/Panitnun-6243/duckduck-server/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SwitchStatusRoutes(app *fiber.App) {
	app.Get("/api/v1/switch-status", middlewares.Jwt(), getSwitchStatusHandler)
	app.Patch("/api/v1/switch-status/:id", middlewares.Jwt(), updateSwitchStatusHandler)
}

func getSwitchStatusHandler(c *fiber.Ctx) error {
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	on, err := services.GetSwitchStatusByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("Cannot get switch status", err))
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(bson.M{"on": on}))
}

func updateSwitchStatusHandler(c *fiber.Ctx) error {
	var requestBody struct {
		SwitchStatus bool `json:"on"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	switchID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid switch ID", err))
	}

	// Extract userID from JWT claims
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	// Check if the user is authorized to update this light control
	_, err = services.GetSwitchStatusByIDAndUserID(switchID, userID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Unauthorized", nil))
	}

	err = services.UpdateUserSwitchStatus(switchID, requestBody.SwitchStatus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Switch status update failed", err))
	}

	// Publish the update to MQTT
	deviceCode := "SSAC12"
	mqttTopic := fmt.Sprintf("%s/power", deviceCode)
	filter := bson.M{
		"on": requestBody.SwitchStatus,
	}
	payload, _ := json.Marshal(filter) // Convert the updatedControl struct to JSON
	client := util.CreateMqttClient()
	util.Publish(client, mqttTopic, string(payload))
	return c.Status(fiber.StatusOK).JSON(responses.Info("Switch status updated successfully"))
}
