package handlers

import (
	"fmt"

	"github.com/donairl/gofiber-dontemplate/lib"
	"github.com/gofiber/fiber/v2"
)

func Dashboard(c *fiber.Ctx) error {
	sess, err := lib.Store.Get(c)
	if err != nil {
		panic(err)
	}

	username := sess.Get("username")
	AuthorizedMessage := fmt.Sprintf("Welcome %v", username)

	csrfToken, ok := c.Locals("csrf").(string)

	return c.Render("dashboard", fiber.Map{
		"Title":   "Dashboard",
		"csrf":    csrfToken,
		"status":  ok,
		"message": AuthorizedMessage,
	}, "layouts/main")
}
func LoginView(c *fiber.Ctx) error {
	// Render index
	csrfToken, ok := c.Locals("csrf").(string)

	return c.Render("login", fiber.Map{
		"Title":  "Login Page",
		"csrf":   csrfToken,
		"status": ok,
	})
}

func RegisterView(c *fiber.Ctx) error {
	// Render register view
	csrfToken, ok := c.Locals("csrf").(string)

	return c.Render("register", fiber.Map{
		"Title":  "Login Page",
		"csrf":   csrfToken,
		"status": ok,
	})
}

func AboutHandler(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title": "Login Page",
	})
}
