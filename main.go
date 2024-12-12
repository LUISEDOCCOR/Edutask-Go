package main

import (
	pages_routes "edutask/src/routes/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./src/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/images", "./src/static/images")
	app.Static("/styles", "./src/static/styles")

	//healthcheck
	app.Get("/ok", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	//pages
	pagesGroup := app.Group("/")
	pages_routes.Routes(pagesGroup)

	app.Listen(":3000")
}
