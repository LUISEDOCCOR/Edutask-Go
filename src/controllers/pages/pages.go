package pages

import (
	waitingList_model "edutask/src/models/redis/waitingList"

	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	counter, err := waitingList_model.GetEmailsCounter()
	if counter == 0 || err != nil {
		counter = 1
	}
	return c.Render("pages/home", fiber.Map{"counter": counter}, "layouts/home_layout")
}
