package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func New() *fiber.App {
	engine := html.New("./views", ".html")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".html"))

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c *fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("login", fiber.Map{
			"Title": "Login Page",
		}, "layouts/main")
	})

	app.Static(
		"/static",  // mount address
		"./public", // path to the file folder
	)

	return app
}
