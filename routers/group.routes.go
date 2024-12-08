package routers

import (
	"github.com/donairl/gofiber-dontemplate/handlers"
	"github.com/donairl/gofiber-dontemplate/repository"
	"github.com/donairl/gofiber-dontemplate/services"
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
	r.Post("edit", handlers.UserEdit)

}

func ProductRoutes(app fiber.Router) {
	r := app.Group("/products")
	productRepo := repository.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	r.Get("/", productHandler.ProductView)
	r.Post("/create", productHandler.ProductCreate)
}
