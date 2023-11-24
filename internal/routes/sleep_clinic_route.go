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

func SleepClinicRoutes(app *fiber.App) {
	app.Get("/api/v1/sleep-clinic", middlewares.Jwt(), getSleepClinicHandler)
	app.Put("/api/v1/sleep-clinic/:id", middlewares.Jwt(), updateSleepClinicHandler)
	// Custom lullaby song routes
	app.Post("/api/v1/custom-lullaby-song", middlewares.Jwt(), addCustomLullabySongHandler)
	app.Get("/api/v1/custom-lullaby-song", middlewares.Jwt(), getCustomLullabySongsHandler)
	app.Get("/api/v1/preset-lullaby-song", getPresetLullabySongsHandler)
}

// Get Sleep Clinic data handler
func getSleepClinicHandler(c *fiber.Ctx) error {
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))
	sleepClinic, err := services.GetSleepClinicByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("Sleep clinic data not found", err))
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(sleepClinic))
}

// Update Sleep Clinic data handler
func updateSleepClinicHandler(c *fiber.Ctx) error {
	var updatedSleepClinic models.SleepClinic
	if err := c.BodyParser(&updatedSleepClinic); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	sleepClinicID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Invalid sleep clinic ID", err))
	}

	// Extract userID from JWT claims
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	// Check if the user is authorized to update this sleep clinic data
	sleepClinic, err := services.GetSleepClinicByUser(userID)
	if err != nil || sleepClinic == nil || sleepClinic.ID != sleepClinicID {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Unauthorized", nil))
	}

	err = services.UpdateUserSleepClinic(sleepClinicID, &updatedSleepClinic)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Sleep clinic update failed", err))
	}

	return c.Status(fiber.StatusOK).JSON(responses.Info("Sleep clinic updated successfully"))
}

func addCustomLullabySongHandler(c *fiber.Ctx) error {
	var input models.LullabyDetail
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	if err := services.AddCustomLullabySongService(userID, input.Name, input.Path, input.Category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Custom lullaby song added successfully"})
}

func getCustomLullabySongsHandler(c *fiber.Ctx) error {
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID, _ := primitive.ObjectIDFromHex(claims["sub"].(string))

	songs, err := services.GetCustomLullabySongsService(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(songs))
}

func getPresetLullabySongsHandler(c *fiber.Ctx) error {
	songs, err := services.GetPresetLullabySongsService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(songs))
}
