package middlewares

import (
	"github.com/Panitnun-6243/duckduck-server/internal/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Jwt() fiber.Handler {
	cfg := config.LoadConfig()
	conf := jwtware.Config{
		SigningKey:   []byte(cfg.JWTSecret),
		TokenLookup:  "header:Authorization",
		AuthScheme:   "Bearer",
		ContextKey:   "l",
		ErrorHandler: jwtError,
	}

	return jwtware.New(conf)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
