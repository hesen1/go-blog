package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func phone(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString("^1[3456789]\\d{9}$", fl.Field().String())
	if !match {
		return false
	}

	return true
}
