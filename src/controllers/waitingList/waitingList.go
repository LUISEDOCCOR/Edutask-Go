package waitinglist_controller

import (
	waitingList_model "edutask/src/models/redis/waitingList"

	"github.com/gofiber/fiber/v2"
)

func AddEmail(c *fiber.Ctx) error {
	email := c.FormValue("email")
	if len(email) > 5 {
		_ = waitingList_model.AddEmail(email)
	}

	return c.Redirect("/")
}
