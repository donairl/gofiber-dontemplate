package handlers

import "github.com/gofiber/fiber/v2"

func Authlogin(c *fiber.Ctx) error {
	return c.SendString("You are login 👋!")
}
