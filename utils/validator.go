package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateInput(data interface{}) ([]string, error) {
	// Create new validation and check the struct
	validate := validator.New()
	var errArr []string
	err := validate.Struct(data)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range err.(validator.ValidationErrors) {
				var message string
				switch validationErr.Tag() {
				case "required":
					message = fmt.Sprintf("%s is required", validationErr.Field())
				case "email":
					message = "Please enter a valid email format"
				case "len":
					message = fmt.Sprintf("%s length invalid", validationErr.Field())
				default:
					message = fmt.Sprintf("%s is invalid %s", validationErr.Field(), validationErr.Tag())
				}
				errArr = append(errArr, message)
			}
		} else {
			// Error lainnya
			errArr = append(errArr, err.Error())
			return errArr, err
		}
		return errArr, err
	}
	return errArr, nil
}
