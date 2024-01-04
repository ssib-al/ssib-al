package shorten

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"ssib.al/ssib-al-back/config"
	"ssib.al/ssib-al-back/ent/link"
	"ssib.al/ssib-al-back/utils/logger"
	"ssib.al/ssib-al-back/utils/randstring"
	"ssib.al/ssib-al-back/utils/validator"
)

type ShortenReq struct {
	Domain   string `json:"domain" validate:"required"`
	Uri      string `json:"uri"`
	Target   string `json:"target" validate:"required"`
	Password string `json:"password"`
}

func ShortenCtrl(c *fiber.Ctx) error {
	body := new(ShortenReq)
	if errArr := validator.ParseAndValidate(c, body); errArr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errArr)
	}
	client := config.GetDBClient()
	uri := body.Uri
	// add database logic here
	if body.Uri == "" {
		noDuplicate := false
		for !noDuplicate {
			// generate random uri
			uri = randstring.Generate(8)
			noDuplicate = !FindDuplicatedLink(body.Domain, uri)
		}
	} else {
		if FindDuplicatedLink(body.Domain, uri) {
			return c.Status(fiber.StatusBadRequest).JSON(config.RaiseError(
				fiber.StatusBadRequest,
				"URI already exists",
				"이미 존재하는 URI입니다.",
			))
		}
	}

	newLinkEntity := client.Link.Create().
		SetDomain(body.Domain).
		SetURI(uri).
		SetTargetURL(body.Target)
	if body.Password != "" {
		logger.Debug("password is not empty")
		// TODO: create user and set owner
	}
	if _, err := newLinkEntity.Save(context.Background()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)

	}
	return c.Status(fiber.StatusOK).JSON(newLinkEntity)
}

func FindDuplicatedLink(domain string, uri string) bool {
	dbClient := config.GetDBClient()

	// check if domain + uri already exists
	_, err := dbClient.Link.
		Query().
		Where(
			link.DomainEQ(domain),
			link.URIEQ(uri),
		).
		First(context.Background())
	return err == nil

}
