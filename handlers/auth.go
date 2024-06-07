package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/donairl/gofiber-dontemplate/lib"
	"github.com/donairl/gofiber-dontemplate/models"
	"github.com/gofiber/fiber/v2"
)

func AuthRegister(c *fiber.Ctx) error {
	user := &models.User{}

	user.Email = c.FormValue("username")
	password := c.FormValue("password")
	repassword := c.FormValue("repassword")
	if password != repassword {
		return c.Status(fiber.StatusBadRequest).SendString("Password not match")
	}

	user.PasswordHash = lib.GeneratePassword(password)

	role, err := strconv.ParseUint(c.FormValue("role"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid role")
	}
	user.Role = uint(role)

	birthdayStr := c.FormValue("birthday")
	println("username: ", user.Email)
	println("Birthday", birthdayStr)
	birthday, err := time.Parse("01-02-2006", birthdayStr)
	if err != nil {
		// handle error if the date format is incorrect
		//return c.Status(fiber.StatusBadRequest).SendString("Invalid date format")
		sess, err := lib.Store.Get(c)
		if err != nil {

			panic(err)
		}
		sess.Set("flash-error", "Invalid date format")
		return c.Redirect(
			"/register",
			http.StatusMovedPermanently,
		)
	}
	user.Birthday = &birthday
	models.UserCreate(user)

	csrfToken, ok := c.Locals("csrf").(string)
	if !ok {

		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Render("login", fiber.Map{
		"Title":   "Login",
		"csrf":    csrfToken,
		"success": "Save success",
	})
}

func AuthLogin(c *fiber.Ctx) error {

	sess, err := lib.Store.Get(c)
	if err != nil {

		panic(err)
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	user := models.UserFindByEmail(username)
	okpass := lib.VerifyPassword(user.PasswordHash, password) == nil

	if okpass {

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

}

func Logout(c *fiber.Ctx) error { // Check if the user is authenticated
	if !IsAuthenticated(c) {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Redirect("/login", fiber.StatusTemporaryRedirect) // Logout the user

}

func IsAuthenticated(c *fiber.Ctx) bool {
	sess, err := lib.Store.Get(c)
	if err != nil {
		return false
	}
	// Check if the user is authenticated
	if sess.Get("name") == "" {
		return false
	}
	return true
}
