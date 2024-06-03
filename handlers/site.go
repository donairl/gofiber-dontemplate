package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Dashboard(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
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
	}, "layouts/main")
}

func AboutHandler(c *fiber.Ctx) error {
	return c.SendString("This is about page")
}
