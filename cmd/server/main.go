package main

import (
	"context"
	"github.com/Panitnun-6243/duckduck-server/internal/config"
	"github.com/Panitnun-6243/duckduck-server/internal/db"
	"github.com/Panitnun-6243/duckduck-server/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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
	app := fiber.New()
	app.Use(cors.New())
	routes.UserRoutes(app)

	log.Fatal(app.Listen(cfg.ServerAddress))
}
