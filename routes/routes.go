package routes

import (
	"github.com/gofiber/fiber/v3"
)

func New() *fiber.App {
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})
	return app
}
