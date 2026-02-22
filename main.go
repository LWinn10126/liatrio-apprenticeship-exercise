package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New() // Create a new Fiber application

	//data := { "message": "Logan Winn","timestamp": ""}

	app.Get("/", func(c fiber.Ctx) error { // Define a route for the root URL  c fiber.Ctx is the request context for headers, parameters, etc.
		err := c.SendString("Hello, World!") // Send a simple string response
		return err
	})

	app.Listen(":3000") // Start the server on port 3000
}
