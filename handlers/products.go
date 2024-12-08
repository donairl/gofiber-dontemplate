package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/donairl/gofiber-dontemplate/lib"
	"github.com/donairl/gofiber-dontemplate/models"
	"github.com/donairl/gofiber-dontemplate/services"
	"github.com/gofiber/fiber/v2"
)

// product view handler
type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) ProductView(c *fiber.Ctx) error {
	sess, err := lib.Store.Get(c)
	if err != nil {
		panic(err)
	}
	//csrf
	csrfToken, _ := c.Locals("csrf").(string)

	if !IsAuthenticated(c) {
		sess.Set("flash-error", "Forbidden, please login first")
		return c.Redirect(
			"/login",
			http.StatusMovedPermanently,
		)
	}

	products, err := h.productService.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	log.Println(products)

	return c.Render("products/productlist", fiber.Map{
		"Title":    "List Product Page",
		"Products": products,
		"csrf":     csrfToken,
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
			return c.Redirect("/product/create", http.StatusSeeOther)
		}

		if err := models.ProductCreate(product); err != nil {
			sess.Set("flash-error", "Failed to create product: "+err.Error())
			return c.Redirect("/product/create", http.StatusSeeOther)
		}

		sess.Set("flash-success", "Product created successfully")
		return c.Redirect("/product/view", http.StatusSeeOther)
	}

	// Handle GET request (display form)
	return c.Render("products/productform", fiber.Map{
		"Title": "Create New Product",
		"csrf":  csrfToken,
	})

}

// product delete handler
func ProductDelete(c *fiber.Ctx) error {
	sess, err := lib.Store.Get(c)
	if err != nil {
		panic(err)
	}

	if !IsAuthenticated(c) {
		sess.Set("flash-error", "Forbidden, please login first")
		return c.Redirect("/login", http.StatusMovedPermanently)
	}

	productID := c.Params("id")
	if productID == "" {
		sess.Set("flash-error", "Product ID is required")
		return c.Redirect("/product/view", http.StatusSeeOther)
	}

	productIDUint, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		sess.Set("flash-error", "Invalid product ID")
		return c.Redirect("/product/view", http.StatusSeeOther)
	}

	if err := models.ProductDelete(uint(productIDUint)); err != nil {
		sess.Set("flash-error", "Failed to delete product: "+err.Error())
		return c.Redirect("/product/view", http.StatusSeeOther)
	}

	sess.Set("flash-success", "Product deleted successfully")
	return c.Redirect("/product/view", http.StatusSeeOther)

}

// product update handler
func ProductUpdate(c *fiber.Ctx) error {
	sess, err := lib.Store.Get(c)
	if err != nil {
		panic(err)
	}

	csrfToken, _ := c.Locals("csrf").(string)

	if !IsAuthenticated(c) {
		sess.Set("flash-error", "Forbidden, please login first")
		return c.Redirect("/login", http.StatusMovedPermanently)
	}

	return c.SendString("You are test")
	//return c.JSON(Products)
}
