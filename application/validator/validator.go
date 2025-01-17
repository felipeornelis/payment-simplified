package validator

import "github.com/go-playground/validator/v10"

func Validate(data interface{}) error {
	return validator.New(validator.WithRequiredStructEnabled()).Struct(data)
}
