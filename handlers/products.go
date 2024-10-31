package handlers

import (
	"log"
	"net/http"

	"github.com/donairl/gofiber-dontemplate/lib"
	"github.com/donairl/gofiber-dontemplate/models"
	"github.com/gofiber/fiber/v2"
)

func ProductView(c *fiber.Ctx) error {
	sess, err := lib.Store.Get(c)
	if err != nil {

		panic(err)
	}

	if !IsAuthenticated(c) {
		sess.Set("flash-error", "Forbidden, please login first")
		return c.Redirect(
			"/login",
			http.StatusMovedPermanently,
		)
	}
	products, err := models.ProductsFindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	log.Println(products)

	return c.Render("products/productlist", fiber.Map{
		"Title":    "List Product Page",
		"Products": products,
	})
}

func ProductCreate(c *fiber.Ctx) error {
	sess, err := lib.Store.Get(c)
	if err != nil {
		panic(err)
	}
	//csrf
	csrfToken, _ := c.Locals("csrf").(string)

	if !IsAuthenticated(c) {
		sess.Set("flash-error", "Forbidden, please login first")
		return c.Redirect("/login", http.StatusMovedPermanently)
	}

	// Handle POST request (form submission)
	if c.Method() == "POST" {
		product := new(models.Product)

		if err := c.BodyParser(product); err != nil {
			sess.Set("flash-error", "Invalid input data")
			return c.Redirect("/products/create", http.StatusSeeOther)
		}

		if err := models.CreateProduct(product); err != nil {
			sess.Set("flash-error", "Failed to create product: "+err.Error())
			return c.Redirect("/products/create", http.StatusSeeOther)
		}

		sess.Set("flash-success", "Product created successfully")
		return c.Redirect("/products", http.StatusSeeOther)
	}

	// Handle GET request (display form)
	return c.Render("products/productform", fiber.Map{
		"Title": "Create New Product",
		"csrf":  csrfToken,
	})
}
