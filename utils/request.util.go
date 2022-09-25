package utils

import (
	"github.com/go-playground/validator"
)

type requestValidator struct {
	validator *validator.Validate
}

func (rv *requestValidator) Validate(i interface{}) error {
	if err := rv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewRequestValidator(validator *validator.Validate) *requestValidator {
	return &requestValidator{validator}
}
