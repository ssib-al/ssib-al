package shorten

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"

	"github.com/gofiber/fiber/v2"
	"ssib.al/ssib-al-back/config"
	"ssib.al/ssib-al-back/constants"
	"ssib.al/ssib-al-back/ent/link"
	"ssib.al/ssib-al-back/utils/crypt"
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
		errJson, _ := json.Marshal(errArr)
		return c.Status(fiber.StatusBadRequest).JSON(config.RaiseError(
			fiber.StatusBadRequest,
			"Invalid request body",
			string(errJson),
		))
	}
	if !slices.Contains(constants.Domain, body.Domain) {
		fmt.Println(constants.Domain)
		return c.Status(fiber.StatusBadRequest).JSON(config.RaiseError(
			fiber.StatusBadRequest,
			"Invalid domain",
			"존재하지 않는 도메인입니다.",
		))
	}
	if body.Password != "" {
		// password validation
		if !validator.ValidatePassword(body.Password) {
			return c.Status(fiber.StatusBadRequest).JSON(config.RaiseError(
				fiber.StatusBadRequest,
				"Invalid password",
				"비밀번호 형식에 맞지 않습니다.",
			))
		}
	}
	client := config.GetDBClient()
	uri := body.Uri
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

	newLinkEntity, err := client.Link.Create().
		SetDomain(body.Domain).
		SetURI(uri).
		SetTargetURL(body.Target).
		Save(context.Background())
	if err != nil {
		logger.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(config.RaiseError(
			fiber.StatusInternalServerError,
			"Internal server error",
			"서버 오류가 발생했습니다.",
		))
	}
	if body.Password != "" {

		newUserEntity, err := client.User.Create().
			SetUsername(body.Domain + "/" + uri).
			SetPasswordHash(crypt.SHA512(body.Password)).
			Save(context.Background())
		if err != nil {
			logger.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(config.RaiseError(
				fiber.StatusInternalServerError,
				"Internal server error",
				"서버 오류가 발생했습니다.",
			))
		}
		_, err = newLinkEntity.Update().
			SetOwner(newUserEntity).
			Save(context.Background())
		if err != nil {
			logger.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(config.RaiseError(
				fiber.StatusInternalServerError,
				"Internal server error",
				"서버 오류가 발생했습니다.",
			))
		}
	}
	return c.Status(fiber.StatusOK).JSON("https://" + body.Domain + "/l/" + uri)
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
