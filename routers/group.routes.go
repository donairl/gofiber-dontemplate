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

func UserRoutes(app fiber.Router) {
	r := app.Group("/user")
	r.Get("/view", handlers.UserCrudView)
	// Generate route for GET with parameter id
	r.Get("edit/:id", handlers.UserEdit)
	r.Delete("/delete/:id", handlers.UserDeleteHandler)

}

func BlogRoutes(app fiber.Router) {
	// r := app.Group("/blog")
	// r.Get("/list", handlers.BlogListHandler)
	// r.Post("/create", handlers.BlogCreateHandler)
	// r.Put("/update/:id", handlers.BlogUpdateHandler)
	// r.Delete("/delete/:id", handlers.BlogDeleteHandler)
}
