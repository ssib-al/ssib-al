package recaptcha

import (
	"github.com/go-resty/resty/v2"
	"ssib.al/ssib-al-back/config"
	"ssib.al/ssib-al-back/utils/logger"
)

type RecaptchaReqBody struct {
	Response string `json:"response"`
}

type RecaptchaResBody struct {
	Success bool `json:"success"`
}

func IsValid(token string) bool {
	client := resty.New()
	resp, err := client.R().
		SetFormData(map[string]string{
			"secret":   config.Env.ReCaptchaSecret,
			"response": token,
		}).
		SetResult(&RecaptchaResBody{}).
		Post("https://www.google.com/recaptcha/api/siteverify")
	if err != nil {
		logger.Error(err)
		return false
	}
	result := resp.Result().(*RecaptchaResBody)
	return result.Success
}
