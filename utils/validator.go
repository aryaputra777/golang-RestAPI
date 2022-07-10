package utils

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func NewValidator() *validator.Validate {

	validate := validator.New()
	return validate
}

func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
