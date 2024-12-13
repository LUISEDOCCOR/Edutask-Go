package main

import (
	"edutask/src/database"
	pages_routes "edutask/src/routes/pages"
	waitinglist_routes "edutask/src/routes/waitingList"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	//database
	database.SetupRedis()

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

	//groups
	rootGroup := app.Group("/")
	apiGroup := app.Group("/api")

	//pages
	pages_routes.Routes(rootGroup)

	waitingListGroup := apiGroup.Group("/waitinglist")
	waitinglist_routes.Router(waitingListGroup)

	app.Listen(":3000")
}
