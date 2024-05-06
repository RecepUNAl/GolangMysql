package helper

import (
	"fmt"
	"github.com/go-playground/validator"
)

type ValidationErrors map[string]string

func TranslateValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s alanı boş olamaz", err.Field())
	case "min":
		return fmt.Sprintf("%s alanı en az %s karakter olmalıdır", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s alanı en fazla %s karakter olmalıdır", err.Field(), err.Param())
	case "email":
		return fmt.Sprintf("%s alanı geçerli bir e-posta adresi olmalı", err.Field())
	default:
		return fmt.Sprintf("%s geçerli değil", err.Field())
	}
}

var Validate = validator.New()

func ValidateStruct(u interface{}) (interface{}, error) {
	err := Validate.Struct(u)

	if err != nil {
		validationErrors := make(ValidationErrors)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = TranslateValidationError(err)
		}
		return validationErrors, err
	}
	return u, nil
}
