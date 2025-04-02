package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Main function to start the application
	// Initialize the application, set up routes, and start the server
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":4444")
}
