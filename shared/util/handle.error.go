package util

import (
	"github.com/go-playground/validator/v10"
)

type ApiErrorValidator struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func HandleErrorValidator(ve validator.ValidationErrors) []ApiErrorValidator {
	out := make([]ApiErrorValidator, len(ve))
	for i, fe := range ve {
		out[i] = ApiErrorValidator{fe.Field(), messageForTag(fe.Tag())}

	}
	return out
}

func messageForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	}
	return ""
}
