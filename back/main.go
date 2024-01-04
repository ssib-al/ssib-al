package main

import (
	"github.com/gofiber/fiber/v2"
	"ssib.al/ssib-al-back/config"
	"ssib.al/ssib-al-back/routers"
	"ssib.al/ssib-al-back/utils/logger"
)

func main() {
	logger.NewLogger()

	config.NewEnvConfig()
	config.NewDBClient()
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routers.SetupRouter(app)

	//server public folder
	app.Static("/assets", "./public/assets")
	app.Static("/font", "./public/font")
	app.Static("/ssibal-logo.svg", "./public/ssibal-logo.svg")
	app.Static("*", "./public/index.html")

	app.Listen(":3000")
}
