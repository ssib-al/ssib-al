package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	//server public folder
	app.Static("/assets", "./public/assets")
	app.Static("/font", "./public/font")
	app.Static("/ssibal-logo.svg", "./public/ssibal-logo.svg")
	app.Static("*", "./public/index.html")

	app.Listen(":3000")
}
