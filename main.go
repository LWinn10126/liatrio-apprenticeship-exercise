package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New() // Create a new Fiber application

	//data := { "message": "Logan Winn","timestamp": ""}

	app.Get("/", func(c fiber.Ctx) error { // Define a route for the root URL  c fiber.Ctx is the request context for headers, parameters, etc.
		err := c.JSON(fiber.Map{"message": "My name is Logan Winn", "timestamp": time.Now().Format("2006-01-02 15:04:05")}) // Send a simple string response
		return err
	})

	// Get port from environment (Cloud Run sets this automatically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local development
	}

	app.Listen(":" + port)
}
