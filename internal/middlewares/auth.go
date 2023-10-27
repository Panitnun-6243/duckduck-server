package middlewares

import (
	"github.com/Panitnun-6243/duckduck-server/internal/config"
	"github.com/Panitnun-6243/duckduck-server/internal/responses"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Jwt() fiber.Handler {
	cfg := config.LoadConfig()
	conf := jwtware.Config{
		SigningKey:  []byte(cfg.JWTSecret),
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		ContextKey:  "l",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return responses.Error("JWT validation failure", err)
		},
	}

	return jwtware.New(conf)
}
