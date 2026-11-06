package api_validator

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func Set() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func Get() *validator.Validate {
	if validate == nil {
		panic("validator should be instanciated")
	}
	return validate
}
