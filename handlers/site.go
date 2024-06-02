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

	return c.SendString("You are test :" + AuthorizedMessage)
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

// func CsrfErrorHandler(c *fiber.Ctx) error {
// 	return c.SendString("Error")
// }
