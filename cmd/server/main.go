package main

import (
	"context"
	"github.com/Panitnun-6243/duckduck-server/config"
	"github.com/Panitnun-6243/duckduck-server/database"
	"github.com/Panitnun-6243/duckduck-server/internal/routes"
	"github.com/Panitnun-6243/duckduck-server/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	db.Connect()
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(db.GetDB().Client(), nil)

	util.CreateMqttClient()
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Welcome to DuckDuck API"})
	})
	routes.UserRoutes(app)
	routes.AlarmRoutes(app)
	routes.LightControlRoutes(app)
	routes.DashboardRoutes(app)
	routes.ConnectionRoutes(app)
	routes.SleepClinicRoutes(app)
	routes.HslLightRoutes(app)
	routes.SwitchStatusRoutes(app)
	routes.CctLightRoutes(app)
	routes.DeviceRoutes(app)

	// Start the server
	log.Fatal(app.Listen(cfg.ServerAddress))
}
