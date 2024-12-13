package waitinglist_routes

import (
	waitinglist_controller "edutask/src/controllers/waitingList"

	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Post("/", waitinglist_controller.AddEmail)
}
