package models

import "github.com/go-playground/validator/v10"

type FieldErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetFieldErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "eqfield":
		return "Fields doesn't match"
	case "email":
		return "Field value must be valid email"
	}

	return "Unkown error"
}
