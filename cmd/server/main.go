package main

import (
	"github.com/Panitnun-6243/duckduck-server/internal/config"
	"github.com/Panitnun-6243/duckduck-server/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitializeDB()
	defer config.DisconnectDB()
	// Create new Fiber app
	app := fiber.New()

	// Set up routes
	routes.SetupRoutes(app)

	// Start the server on port 3000
	err := app.Listen(":5050")
	if err != nil {
		return
	}

}
