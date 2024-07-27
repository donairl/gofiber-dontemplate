package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/donairl/gofiber-dontemplate/lib"
	"github.com/donairl/gofiber-dontemplate/models"
	"github.com/gofiber/fiber/v2"
)

func Usertest(c *fiber.Ctx) error {

	return c.SendString("You are test")
}

// render user view
func UserCrudView(c *fiber.Ctx) error {

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

	Users := models.UserFindAll()
	log.Println(Users)
	return c.Render("userlist", fiber.Map{
		"Title": "List User Page",
		"Users": Users,
	})

}

// user delete handler
func UserDeleteHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Invalid user ID")
	}
	errdel := models.UserDelete(uint(userID))

	if errdel != nil {
		return c.Status(500).SendString("Failed to delete user")
	}
	return c.Redirect("/user/list", fiber.StatusSeeOther)
}

// user delete handler
func UserEdit(c *fiber.Ctx) error {

	csrfToken, ok := c.Locals("csrf").(string)
	sess, err := lib.Store.Get(c)
	if err != nil {
		panic(err)
	}

	if c.Context().IsPost() {
		//get from formData
		user := &models.User{}

		user.Email = c.FormValue("username")
		password := c.FormValue("password")
		repassword := c.FormValue("repassword")

		if password != "" || repassword != "" {
			if password != repassword {
				return c.Status(fiber.StatusBadRequest).SendString("Password not match")
			}

			user.PasswordHash = lib.GeneratePassword(password)
		}

		role, err := strconv.ParseUint(c.FormValue("role"), 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid role")
		}

		user.Role = uint(role)

		birthdayStr := c.FormValue("birthday")
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
		user.Fullname = c.FormValue("fullname")
		models.UserUpdate(user)

	} else {

		id := c.Params("id")
		userID, err := strconv.ParseUint(id, 10, 32)

		if err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Invalid user ID")
		}

		flashError := sess.Get("flash-error")
		usermodel := models.UserFindByID(uint(userID))
		log.Println(usermodel)

		return c.Render("edit-user", fiber.Map{
			"Title":  "Edit User",
			"csrf":   csrfToken,
			"status": ok,
			"error":  flashError,
			"User":   usermodel,
		}, "layouts/blank")

	}

	return c.Redirect("/user/view", fiber.StatusSeeOther)
}
