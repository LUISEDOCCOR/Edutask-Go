package pages_routes

import (
	"edutask/src/controllers/pages"

	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	router.Get("/", pages.HomePage)
}
