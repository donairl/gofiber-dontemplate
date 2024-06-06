package handlers

import (
	"net/http"
	"time"

	"github.com/donairl/gofiber-dontemplate/lib"
	"github.com/donairl/gofiber-dontemplate/models"
	"github.com/gofiber/fiber/v2"
)

func AuthRegister(c *fiber.Ctx) error {
	user := &models.User{}

	user.Email = c.FormValue("username")
	user.PasswordHash = c.FormValue("password")
	//	repassword := c.FormValue("repassword")
	user.Fullname = c.FormValue("fullname")
	user.Role = 1

	birthdayStr := c.FormValue("birthday")
	println("username: ", user.Email)
	println("Birthday", birthdayStr)
	birthday, err := time.Parse("2006-01-02", birthdayStr)
	if err != nil {
		// handle error if the date format is incorrect
		return c.Status(fiber.StatusBadRequest).SendString("Invalid date format")
	}
	user.Birthday = &birthday
	models.UserCreate(user)

	csrfToken, ok := c.Locals("csrf").(string)
	if !ok {

		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Render("login", fiber.Map{
		"Title": "Login",
		"csrf":  csrfToken,
		"error": "Save success",
	})
}

func AuthLogin(c *fiber.Ctx) error {

	sess, err := lib.Store.Get(c)
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
