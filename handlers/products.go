package handlers
import (
	"github.com/gofiber/fiber/v2"
)

func Producttest(c *fiber.Ctx) error {

	return c.SendString("You are test")
}