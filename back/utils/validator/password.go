package validator

import "regexp"

const (
	regexpPassword = `^(?:(?:.*[a-zA-Z])(?:.*\d)(?:.*[!@#$%^&*()-=_+]).+)$`
)

func ValidatePassword(password string) bool {
	if len(password) < 8 || len(password) > 32 {
		return false
	}

	passwordRegex := regexp.MustCompile(regexpPassword)
	return passwordRegex.MatchString(password)
}
