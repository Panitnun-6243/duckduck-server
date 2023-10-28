package routes

import (
	"github.com/Panitnun-6243/duckduck-server/internal/middlewares"
	"github.com/Panitnun-6243/duckduck-server/internal/models"
	"github.com/Panitnun-6243/duckduck-server/internal/responses"
	"github.com/Panitnun-6243/duckduck-server/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func UserRoutes(app *fiber.App) {
	app.Post("/api/v1/register", registerHandler)
	app.Post("/api/v1/login", loginHandler)
	app.Get("/api/v1/users", middlewares.Jwt(), getUserInfoHandler)
	app.Patch("/api/v1/users", middlewares.Jwt(), updateUserProfileHandler)
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

func getUserInfoHandler(c *fiber.Ctx) error {
	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)
	user, err := services.GetUserInfo(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("User not found", err))
	}
	userProfile := models.UserProfile{
		Email:     user.Email,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(userProfile))
}

func updateUserProfileHandler(c *fiber.Ctx) error {
	var payload struct {
		Name      string `json:"name" validate:"required"`
		AvatarURL string `json:"avatar_url" validate:"required,url"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Bad request", err))
	}

	claims := c.Locals("l").(*jwt.Token).Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)
	updatedUser, err := services.UpdateUserProfile(userID, payload.Name, payload.AvatarURL)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("User update failed", err))
	}
	userProfile := models.UserProfile{
		Email:     updatedUser.Email,
		Name:      updatedUser.Name,
		AvatarURL: updatedUser.AvatarURL,
	}
	return c.Status(fiber.StatusOK).JSON(responses.Info(userProfile))
}
