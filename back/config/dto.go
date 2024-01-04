package config

import "github.com/gofiber/fiber/v2"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func RaiseError(code int, message string, detail string) fiber.Map {
	return fiber.Map{
		"error": Error{
			Code:    code,
			Message: message,
			Detail:  detail,
		},
	}
}
