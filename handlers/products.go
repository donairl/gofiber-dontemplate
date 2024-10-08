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

	if !IsAuthenticated(c) {
		sess.Set("flash-error", "Forbidden, please login first")
		return c.Redirect("/login", http.StatusMovedPermanently)
	}

	return c.SendString("You are test")
	//return c.JSON(Products)
}
