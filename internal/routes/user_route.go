package routes

import (
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/responses"
	"github.com/Panitnun-6243/duckduck-server/internal/services"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/api/v1/register", registerHandler)
	app.Post("/api/v1/login", loginHandler)
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

func loginHandler(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	token, err := services.LoginUser(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.Error("Login failed", err))
	}

	return c.Status(fiber.StatusOK).JSON(responses.Info(map[string]string{
		"token": token,
	}))
}
