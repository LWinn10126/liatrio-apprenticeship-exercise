package main

import (
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

	app.Listen(":3000") // Start the server on port 3000
}
