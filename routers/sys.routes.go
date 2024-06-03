package routers

import (
	"github.com/donairl/gofiber-dontemplate/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SysRoutes(app fiber.Router) {
	r := app.Group("/sys")
	r.Get("/info", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	r.Get("/about", handlers.AboutHandler)
}
