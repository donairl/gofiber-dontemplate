package handlers

import (
	"log"
	"strconv"

	"github.com/donairl/gofiber-dontemplate/models"
	"github.com/gofiber/fiber/v2"
)

func Usertest(c *fiber.Ctx) error {

	return c.SendString("You are test")
}

// render user view
func UserCrudView(c *fiber.Ctx) error {

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
