package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func Authtest(c *fiber.Ctx) error {
	return c.SendString("You are great")
}

func Authlogin(c *fiber.Ctx) error {

	sess, err := Store.Get(c)
	if err != nil {

		panic(err)
	}
	username := c.FormValue("username")
	password := c.FormValue("password")

	if (username == "donny.airlangga@gmail.com") && (password == "1234") {

		sess.Set("name", username)
		if err := sess.Save(); err != nil {
			panic(err)
		}
		//return c.SendString("You are login 👋! " + username + " : " + password)
		return c.Redirect(
			"/dashboard",
			http.StatusMovedPermanently,
		)
	} else {

		csrfToken, ok := c.Locals("csrf").(string)
		if !ok {

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Render("login", fiber.Map{
			"Title": "Login",
			"csrf":  csrfToken,
			"error": "Invalid credentials",
		})
	}
	// Check if the credentials are valid
	// user, exists := users[username]
	// var checkPassword string
	// if exists {
	// 	checkPassword = user.Password
	// } else {
	// 	checkPassword = emptyHashString
	// }

	// if bcrypt.CompareHashAndPassword([]byte(checkPassword), []byte(password)) != nil {
	// 	// Authentication failed
	// 	csrfToken, ok := c.Locals("csrf").(string)
	// 	if !ok {
	// 		return c.SendStatus(fiber.StatusInternalServerError)
	// 	}
	//
	// 	return c.Render("login", fiber.Map{
	// 		"Title": "Login",
	// 		"csrf":  csrfToken,
	// 		"error": "Invalid credentials",
	// 	})
	// }

}
