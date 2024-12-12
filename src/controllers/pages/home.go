package pages

import "github.com/gofiber/fiber/v2"

func HomePage(c *fiber.Ctx) error {
	return c.Render("pages/home", nil, "layouts/home_layout")
}
