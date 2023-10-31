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

func CctLightRoutes(app *fiber.App) {
	app.Get("/api/v1/cct-light", middlewares.Jwt(), getCctLightHandler)
	app.Patch("/api/v1/cct-light/:id", middlewares.Jwt(), updateCctLightHandler)
}

func getCctLightHandler(c *fiber.Ctx) error {
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	cctLight, err := services.GetCctLightByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("Cannot get CCT values", err))
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(bson.M{"brightness": cctLight.BrightnessPercentage, "temp": cctLight.CctTemp}))
}

func updateCctLightHandler(c *fiber.Ctx) error {
	var requestBody struct {
		Temp       int     `json:"temp"`
		Brightness float64 `json:"brightness"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	cctID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid cct light ID", err))
	}

	// Extract userID from JWT claims
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	// Check if the user is authorized to update this light control
	_, err = services.GetCctLightByIDAndUserID(cctID, userID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Unauthorized", nil))
	}

	err = services.UpdateUserCctLight(cctID, requestBody.Brightness, requestBody.Temp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Cct light update failed", err))
	}

	// Publish the update to MQTT
	deviceCode := "SSAC12"
	mqttTopic := fmt.Sprintf("%s/cct", deviceCode)
	filter := bson.M{
		"brightness": requestBody.Brightness,
		"temp":       requestBody.Temp,
	}
	payload, _ := json.Marshal(filter)
	client := util.CreateMqttClient()
	util.Publish(client, mqttTopic, string(payload))

	return c.Status(fiber.StatusOK).JSON(responses.Info("Cct light updated successfully"))
}
