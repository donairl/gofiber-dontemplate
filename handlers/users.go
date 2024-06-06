package handlers

import (
	"github.com/donairl/gofiber-dontemplate/models"
	"github.com/gofiber/fiber/v2"
)

func Usertest(c *fiber.Ctx) error {

	return c.SendString("You are test")
}

// render user view
func UserCrudView(c *fiber.Ctx) error {

	Users := models.UserFindAll()
	return c.Render("userlist", fiber.Map{
		"Title": "List User Page",
		"Users": Users,
	})

}
