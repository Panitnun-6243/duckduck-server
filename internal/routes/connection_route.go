package routes

import (
	"github.com/Panitnun-6243/duckduck-server/internal/middlewares"
	"github.com/Panitnun-6243/duckduck-server/internal/responses"
	"github.com/Panitnun-6243/duckduck-server/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConnectionRoutes(app *fiber.App) {
	app.Get("/api/v1/connection-status", middlewares.Jwt(), getConnectionStatusHandler)
	app.Patch("/api/v1/connection-status/:id", middlewares.Jwt(), updateConnectionStatusHandler)
}

func getConnectionStatusHandler(c *fiber.Ctx) error {
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	isConnected, err := services.GetConnectionStatusByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("Cannot get connection status", err))
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(bson.M{"connected": isConnected}))
}

func updateConnectionStatusHandler(c *fiber.Ctx) error {
	var requestBody struct {
		Connected bool `json:"connected"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	connID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid connection ID", err))
	}

	// Extract userID from JWT claims
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	// Check if the user is authorized to update this light control
	_, err = services.GetConnectionStatusByIDAndUserID(connID, userID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Unauthorized", nil))
	}

	err = services.UpdateUserConnectionStatus(connID, requestBody.Connected)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Connection status update failed", err))
	}

	return c.Status(fiber.StatusOK).JSON(responses.Info("Connection status updated successfully"))
}
