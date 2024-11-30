package middleware

import (
	"backend/shared"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type ErrorResponse struct {
	Error       string
	FailedField string
	Tag         string
	Value       interface{}
}

func ValidateRequest(data interface{}) error {
	validate := validator.New()
	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			switch err.Tag() {
			case "email":
				return errors.New(shared.ErrInvalidFieldFormat)
			case "required":
				return errors.New(shared.ErrInvalidFieldFormat)
			case "min":
				return errors.New(shared.ErrInvalidFieldFormat)
			}
		}
	}

	return nil
}
