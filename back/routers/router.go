package routers

import (
	"github.com/gofiber/fiber/v2"
	"ssib.al/ssib-al-back/controllers/link"
)

func SetupRouter(app *fiber.App) {
	linkRouter := app.Group("/l")
	linkRouter.Get("/:uri", func(c *fiber.Ctx) error {
		return link.LinkCtrl(c)
	})
	apiRoter := app.Group("/api")

	apiRoter.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	shortenRouter(apiRoter)
}
