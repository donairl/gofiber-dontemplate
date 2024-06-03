package routers

import (
	"fmt"
	"time"

	"github.com/donairl/gofiber-dontemplate/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func New() *fiber.App {
	engine := html.New("./views", ".html")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".html"))

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// Initialize a session store
	sessConfig := session.Config{
		Expiration:     60 * time.Minute,        // Expire sessions after 30 minutes of inactivity
		KeyLookup:      "cookie:__Host-session", // Recommended to use the __Host- prefix when serving the app over TLS
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
	}

	store := session.New(sessConfig)
	handlers.Store = store

	csrfErrorHandler := func(c *fiber.Ctx, err error) error {
		// Log the error so we can track who is trying to perform CSRF attacks
		// customize this to your needs
		fmt.Printf("CSRF Error: %v Request: %v From: %v\n", err, c.OriginalURL(), c.IP())

		// check accepted content types
		switch c.Accepts("html", "json") {
		case "json":
			// Return a 403 Forbidden response for JSON requests
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "403 Forbidden",
			})
		case "html":
			// Return a 403 Forbidden response for HTML requests
			return c.Status(fiber.StatusForbidden).Render("error", fiber.Map{
				"Title":     "Error",
				"Error":     "403 Forbidden",
				"ErrorCode": "403",
			})
		default:
			// Return a 403 Forbidden response for all other requests
			return c.Status(fiber.StatusForbidden).SendString("403 Forbidden")
		}
	}

	// Configure the CSRF middleware
	csrfConfig := csrf.Config{
		Session:        handlers.Store,
		KeyLookup:      "form:_csrf",  // In this example, we will be using a hidden input field to store the CSRF token
		CookieName:     "__Host-csrf", // Recommended to use the __Host- prefix when serving the app over TLS
		CookieSameSite: "Lax",         // Recommended to set this to Lax or Strict
		CookieSecure:   true,          // Recommended to set to true when serving the app over TLS
		CookieHTTPOnly: true,          // Recommended, otherwise if using JS framework recomend: false and KeyLookup: "header:X-CSRF-Token"
		ContextKey:     "csrf",
		ErrorHandler:   csrfErrorHandler,
		Expiration:     30 * time.Minute,
	}
	csrfMiddleware := csrf.New(csrfConfig)

	app.Use(csrfMiddleware)

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Define a route for the GET method on the root path '/'
	app.Get("/dashboard", handlers.Dashboard)

	app.Get("/", handlers.LoginView)
	app.Get("/login", handlers.LoginView)
	app.Post("/login", handlers.Authlogin)
	app.Get("/dashboard", handlers.Dashboard)

	SysRoutes(app)

	app.Static(
		"/static",  // mount address
		"./public", // path to the file folder
	)

	return app
}
