package routers

import (
	"github.com/gofiber/fiber/v2"
	"ssib.al/ssib-al-back/controllers/shorten"
)

func shortenRouter(apiGroup fiber.Router) {
	apiGroup.Post("/shorten", func(c *fiber.Ctx) error {
		return shorten.ShortenCtrl(c)
	})
}
