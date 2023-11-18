package routes

import (
	"encoding/json"
	"fmt"
	"github.com/Panitnun-6243/duckduck-server/internal/middlewares"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/repositories"
	"github.com/Panitnun-6243/duckduck-server/internal/responses"
	"github.com/Panitnun-6243/duckduck-server/internal/services"
	"github.com/Panitnun-6243/duckduck-server/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AlarmRoutes(app *fiber.App) {
	app.Post("/api/v1/alarms", middlewares.Jwt(), createAlarmHandler)
	app.Get("/api/v1/alarms", middlewares.Jwt(), getAlarmsHandler)
	app.Put("/api/v1/alarms/:id", middlewares.Jwt(), updateAlarmHandler)
	app.Delete("/api/v1/alarms/:id", middlewares.Jwt(), deleteAlarmHandler)
	app.Post("/api/v1/alarms/:alarmId/trigger", middlewares.Jwt(), triggerAlarmHandler)
}

func createAlarmHandler(c *fiber.Ctx) error {
	var alarm models.Alarm
	if err := c.BodyParser(&alarm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	strUserID := claims["sub"].(string)
	userID, err := primitive.ObjectIDFromHex(strUserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid user ID", err))
	}

	alarm.UserID = userID
	createdAlarm, err := services.CreateNewAlarm(&alarm)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Alarm creation failed", err))
	}

	// Publish the update to MQTT
	deviceCode, err := services.GetDeviceCodeByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Failed to get device code", err))
	}
	mqttTopic := fmt.Sprintf("%s/create-alarm", deviceCode)
	payload, _ := json.Marshal(createdAlarm) // Convert the updatedControl struct to JSON
	client := util.CreateMqttClient()
	util.Publish(client, mqttTopic, string(payload))

	return c.Status(fiber.StatusOK).JSON(responses.Info(createdAlarm))
}

func getAlarmsHandler(c *fiber.Ctx) error {
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))
	alarms, err := services.GetAlarmsByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("Alarms not found", err))
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(alarms))
}

func updateAlarmHandler(c *fiber.Ctx) error {
	var updatedAlarm models.Alarm
	if err := c.BodyParser(&updatedAlarm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	alarmID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid alarm ID", err))
	}

	// Check if the user is authorized to update this alarm
	alarm, err := repositories.FindAlarmByAlarmIDAndUserID(alarmID, userID)
	if err != nil || alarm == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Unauthorized", nil))
	}

	err = services.UpdateUserAlarm(alarmID, &updatedAlarm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Alarm update failed", err))
	}

	// Publish the update to MQTT
	deviceCode, err := services.GetDeviceCodeByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Failed to get device code", err))
	}
	mqttTopic := fmt.Sprintf("%s/update-alarm", deviceCode)
	updatedAlarm.ID = alarmID
	updatedAlarm.UserID = userID
	payload, _ := json.Marshal(updatedAlarm) // Convert the updatedControl struct to JSON
	client := util.CreateMqttClient()
	util.Publish(client, mqttTopic, string(payload))

	return c.Status(fiber.StatusOK).JSON(responses.Info("Alarm updated successfully"))
}

func deleteAlarmHandler(c *fiber.Ctx) error {
	// Extract alarmID from request params
	alarmID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid alarm ID", err))
	}

	// Extract userID from JWT claims
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	// Check if the user is authorized to delete this alarm
	alarm, err := repositories.FindAlarmByAlarmIDAndUserID(alarmID, userID)
	if err != nil || alarm == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Unauthorized", nil))
	}

	// If authorized, proceed with the deletion
	err = services.RemoveAlarm(alarmID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("Alarm deletion failed", err))
	}

	// Publish the update to MQTT
	deviceCode, err := services.GetDeviceCodeByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Failed to get device code", err))
	}
	mqttTopic := fmt.Sprintf("%s/delete-alarm", deviceCode)
	filter := bson.M{
		"id": alarmID,
	}
	payload, _ := json.Marshal(filter) // Convert the updatedControl struct to JSON
	client := util.CreateMqttClient()
	util.Publish(client, mqttTopic, string(payload))

	return c.Status(fiber.StatusOK).JSON(responses.Info("Alarm deleted successfully"))
}

func triggerAlarmHandler(c *fiber.Ctx) error {
	var triggerRequest models.TriggerAlarmRequest
	if err := c.BodyParser(&triggerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	alarmID, err := primitive.ObjectIDFromHex(c.Params("alarmId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid alarm ID", err))
	}

	// Extract userID from JWT claims
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	// Check if the user is authorized to trigger this alarm
	alarm, err := repositories.FindAlarmByAlarmIDAndUserID(alarmID, userID)
	if err != nil || alarm == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Unauthorized", nil))
	}

	err = services.TriggerAlarm(userID, alarmID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Failed to trigger alarm", err))
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info("Alarm is triggered successfully"))
}
