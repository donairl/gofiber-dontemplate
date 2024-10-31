package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/donairl/gofiber-dontemplate/lib"
	"github.com/donairl/gofiber-dontemplate/models"
	"github.com/gofiber/fiber/v2"
)

// product view handler
func ProductView(c *fiber.Ctx) error {
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
		"csrf":     csrfToken,
	})
}

// product create handler
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

	if c.Method() == "POST" {
		productID := c.FormValue("product_id")
		if productID == "" {
			sess.Set("flash-error", "Product ID is required")
			return c.Redirect("/product/view", http.StatusSeeOther)
		}

		productIDUint, err := strconv.ParseUint(productID, 10, 32)
		if err != nil {
			sess.Set("flash-error", "Invalid product ID")
			return c.Redirect("/product/view", http.StatusSeeOther)
		}

		product, err := models.GetProductByID(uint(productIDUint))
		if err != nil {
			sess.Set("flash-error", "Product not found")
			return c.Redirect("/product/view", http.StatusSeeOther)
		}

		product.Name = c.FormValue("name")
		product.Description = c.FormValue("description")
		product.Price, _ = strconv.ParseFloat(c.FormValue("price"), 64)
		product.Weight, _ = strconv.ParseFloat(c.FormValue("weight"), 64)
		product.Unit = c.FormValue("unit")
		product.ProductType = c.FormValue("product_type")

		if err := models.ProductUpdate(&product); err != nil {
			sess.Set("flash-error", "Failed to update product: "+err.Error())
			return c.Redirect("/product/view", http.StatusSeeOther)
		}

		sess.Set("flash-success", "Product updated successfully")
		return c.Redirect("/product/view", http.StatusSeeOther)
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

	product, err := models.GetProductByID(uint(productIDUint))
	if err != nil {
		sess.Set("flash-error", "Product not found")
		return c.Redirect("/product/view", http.StatusSeeOther)
	}

	return c.Render("products/productform", fiber.Map{
		"Title":   "Edit Product",
		"Product": product,
		"csrf":    csrfToken,
	})

}
