package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New() // Create a new Fiber application

	app.Get("/", func(c fiber.Ctx) error { // Define a route for the root URL  c fiber.Ctx is the request context for headers, parameters, etc.
		version := os.Getenv("APP_VERSION") // Get the APP_VERSION environment variable

		// Send a simple string response with time allocation and version
		return c.JSON(fiber.Map{"message": "My name is Logan Winn", "timestamp": time.Now().UnixMilli(), "version": version})
	})

	// Get port from environment (Cloud Run sets this automatically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local development
	}

	app.Listen(":" + port)
}
