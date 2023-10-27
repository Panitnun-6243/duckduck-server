package routes

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/responses"
	"github.com/Panitnun-6243/duckduck-server/internal/services"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/api/v1/register", registerHandler)
}

func registerHandler(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	registeredUser, err := services.RegisterUser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Registration failed", err))
	}

	return c.Status(fiber.StatusOK).JSON(responses.Info(registeredUser))
}

//func loginUser(c *fiber.Ctx) error {
//	var request struct {
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}
//
//	if err := c.BodyParser(&request); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Failed to parse request"))
//	}
//
//	token, err := services.LoginUser(request.Email, request.Password)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("Login failed", err))
//	}
//
//	return c.Status(fiber.StatusOK).JSON(responses.Info(token, "Login successful"))
//}
